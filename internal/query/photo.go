package query

import (
	"fmt"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/pkg/capture"
	"github.com/photoprism/photoprism/pkg/rnd"
	"github.com/photoprism/photoprism/pkg/txt"
	"github.com/ulule/deepcopier"
)

// PhotoResult contains found photos and their main file plus other meta data.
type PhotoResult struct {
	// Photo
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	TakenAt          time.Time
	TakenAtLocal     time.Time
	TakenSrc         string
	TimeZone         string
	PhotoUUID        string
	PhotoPath        string
	PhotoName        string
	PhotoTitle       string
	PhotoYear        int
	PhotoMonth       int
	PhotoCountry     string
	PhotoFavorite    bool
	PhotoPrivate     bool
	PhotoLat         float32
	PhotoLng         float32
	PhotoAltitude    int
	PhotoIso         int
	PhotoFocalLength int
	PhotoFNumber     float32
	PhotoExposure    string
	PhotoQuality     int
	PhotoResolution  int
	Merged           bool

	// Camera
	CameraID    uint
	CameraModel string
	CameraMake  string

	// Lens
	LensID    uint
	LensModel string
	LensMake  string

	// Location
	LocationID string
	PlaceID    string
	LocLabel   string
	LocCity    string
	LocState   string
	LocCountry string

	// File
	FileID          uint
	FileUUID        string
	FilePrimary     bool
	FileMissing     bool
	FileName        string
	FileHash        string
	FileType        string
	FileMime        string
	FileWidth       int
	FileHeight      int
	FileOrientation int
	FileAspectRatio float32
	FileColors      string // todo: remove from result?
	FileChroma      uint8  // todo: remove from result?
	FileLuminance   string // todo: remove from result?
	FileDiff        uint32 // todo: remove from result?

	Files []entity.File
}

type PhotoResults []PhotoResult

func (m PhotoResults) Merged() (PhotoResults, int, error) {
	count := len(m)
	merged := make([]PhotoResult, 0, count)

	var lastId uint
	var i int

	for _, res := range m {
		file := entity.File{}

		if err := deepcopier.Copy(&file).From(res); err != nil {
			return merged, count, err
		}

		file.ID = res.FileID

		if lastId == res.ID && i > 0 {
			merged[i-1].Files = append(merged[i-1].Files, file)
			merged[i-1].Merged = true
			continue
		}

		lastId = res.ID

		res.Files = append(res.Files, file)
		merged = append(merged, res)

		i++
	}

	return merged, count, nil
}

func (m *PhotoResult) ShareFileName() string {
	var name string

	if m.PhotoTitle != "" {
		name = strings.Title(slug.MakeLang(m.PhotoTitle, "en"))
	} else {
		name = m.PhotoUUID
	}

	taken := m.TakenAtLocal.Format("20060102-150405")
	token := rnd.Token(3)

	result := fmt.Sprintf("%s-%s-%s.%s", taken, name, token, m.FileType)

	return result
}

// Photos searches for photos based on a Form and returns a PhotoResult slice.
func (q *Query) Photos(f form.PhotoSearch) (results PhotoResults, count int, err error) {
	if err := f.ParseQueryString(); err != nil {
		return results, 0, err
	}

	defer log.Debug(capture.Time(time.Now(), fmt.Sprintf("photos: %+v", f)))

	s := q.db.NewScope(nil).DB()

	// s.LogMode(true)

	s = s.Table("photos").
		Select(`photos.*,
		files.id AS file_id, files.file_uuid, files.file_primary, files.file_missing, files.file_name, files.file_hash, 
		files.file_type, files.file_mime, files.file_width, files.file_height, files.file_aspect_ratio, 
		files.file_orientation, files.file_main_color, files.file_colors, files.file_luminance, files.file_chroma,
		files.file_diff,
		cameras.camera_make, cameras.camera_model,
		lenses.lens_make, lenses.lens_model,
		places.loc_label, places.loc_city, places.loc_state, places.loc_country
		`).
		Joins("JOIN files ON files.photo_id = photos.id AND files.file_type = 'jpg' AND files.file_missing = 0 AND files.deleted_at IS NULL").
		Joins("JOIN cameras ON cameras.id = photos.camera_id").
		Joins("JOIN lenses ON lenses.id = photos.lens_id").
		Joins("JOIN places ON photos.place_id = places.id").
		Joins("LEFT JOIN photos_labels ON photos_labels.photo_id = photos.id AND photos_labels.uncertainty < 100").
		Group("photos.id, files.id")

	if f.ID != "" {
		s = s.Where("photos.photo_uuid = ?", f.ID)
		s = s.Order("files.file_primary DESC")

		if result := s.Scan(&results); result.Error != nil {
			return results, 0, result.Error
		}

		if f.Merged {
			return results.Merged()
		}

		return results, len(results), nil
	}

	var categories []entity.Category
	var label entity.Label
	var labelIds []uint

	if f.Label != "" {
		slugString := strings.ToLower(f.Label)
		if result := q.db.First(&label, "label_slug =? OR custom_slug = ?", slugString, slugString); result.Error != nil {
			log.Errorf("search: label %s not found", txt.Quote(f.Label))
			return results, 0, fmt.Errorf("label %s not found", txt.Quote(f.Label))
		} else {
			labelIds = append(labelIds, label.ID)

			q.db.Where("category_id = ?", label.ID).Find(&categories)

			for _, category := range categories {
				labelIds = append(labelIds, category.LabelID)
			}

			s = s.Where("photos_labels.label_id IN (?)", labelIds)
		}
	}

	if f.Location == true {
		s = s.Where("location_id > 0")

		if f.Query != "" {
			s = s.Joins("LEFT JOIN photos_keywords ON photos_keywords.photo_id = photos.id").
				Joins("LEFT JOIN keywords ON photos_keywords.keyword_id = keywords.id").
				Where("keywords.keyword LIKE ?", strings.ToLower(txt.Clip(f.Query, txt.ClipKeyword))+"%")
		}
	} else if f.Query != "" {
		if len(f.Query) < 2 {
			return results, 0, fmt.Errorf("query too short")
		}

		slugString := slug.Make(f.Query)
		lowerString := strings.ToLower(f.Query)
		likeString := txt.Clip(lowerString, txt.ClipKeyword) + "%"

		s = s.Joins("LEFT JOIN photos_keywords ON photos_keywords.photo_id = photos.id").
			Joins("LEFT JOIN keywords ON photos_keywords.keyword_id = keywords.id")

		if result := q.db.First(&label, "label_slug = ? OR custom_slug = ?", slugString, slugString); result.Error != nil {
			log.Infof("search: label %s not found, using fuzzy search", txt.Quote(f.Query))

			s = s.Where("keywords.keyword LIKE ?", likeString)
		} else {
			labelIds = append(labelIds, label.ID)

			q.db.Where("category_id = ?", label.ID).Find(&categories)

			for _, category := range categories {
				labelIds = append(labelIds, category.LabelID)
			}

			log.Infof("search: label %s includes %d categories", txt.Quote(label.LabelName), len(labelIds))

			s = s.Where("photos_labels.label_id IN (?) OR keywords.keyword LIKE ?", labelIds, likeString)
		}
	}

	if f.Archived {
		s = s.Where("photos.deleted_at IS NOT NULL")
	} else {
		s = s.Where("photos.deleted_at IS NULL")

		if f.Private {
			s = s.Where("photos.photo_private = 1")
		} else if f.Public {
			s = s.Where("photos.photo_private = 0")
		}

		if f.Review {
			s = s.Where("photos.photo_quality < 3")
		} else if f.Quality != 0 && f.Private == false {
			s = s.Where("photos.photo_quality >= ?", f.Quality)
		}
	}

	if f.Error {
		s = s.Where("files.file_error <> ''")
	}

	if f.Album != "" {
		s = s.Joins("JOIN photos_albums ON photos_albums.photo_uuid = photos.photo_uuid").Where("photos_albums.album_uuid = ?", f.Album)
	}

	if f.Camera > 0 {
		s = s.Where("photos.camera_id = ?", f.Camera)
	}

	if f.Lens > 0 {
		s = s.Where("photos.lens_id = ?", f.Lens)
	}

	if f.Year > 0 {
		s = s.Where("photos.photo_year = ?", f.Year)
	}

	if f.Month > 0 {
		s = s.Where("photos.photo_month = ?", f.Month)
	}

	if f.Color != "" {
		s = s.Where("files.file_main_color = ?", strings.ToLower(f.Color))
	}

	if f.Favorites {
		s = s.Where("photos.photo_favorite = 1")
	}

	if f.Story {
		s = s.Where("photos.photo_story = 1")
	}

	if f.Country != "" {
		s = s.Where("photos.photo_country = ?", f.Country)
	}

	if f.Title != "" {
		s = s.Where("LOWER(photos.photo_title) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(f.Title)))
	}

	if f.Hash != "" {
		s = s.Where("files.file_hash = ?", f.Hash)
	}

	if f.Duplicate {
		s = s.Where("files.file_duplicate = 1")
	}

	if f.Portrait {
		s = s.Where("files.file_portrait = 1")
	}

	if f.Mono {
		s = s.Where("files.file_chroma = 0")
	} else if f.Chroma > 9 {
		s = s.Where("files.file_chroma > ?", f.Chroma)
	} else if f.Chroma > 0 {
		s = s.Where("files.file_chroma > 0 AND files.file_chroma <= ?", f.Chroma)
	}

	if f.Diff != 0 {
		s = s.Where("files.file_diff = ?", f.Diff)
	}

	if f.Fmin > 0 {
		s = s.Where("photos.photo_f_number >= ?", f.Fmin)
	}

	if f.Fmax > 0 {
		s = s.Where("photos.photo_f_number <= ?", f.Fmax)
	}

	if f.Dist == 0 {
		f.Dist = 20
	} else if f.Dist > 5000 {
		f.Dist = 5000
	}

	// Inaccurate distance search, but probably 'good enough' for now
	if f.Lat > 0 {
		latMin := f.Lat - SearchRadius*float32(f.Dist)
		latMax := f.Lat + SearchRadius*float32(f.Dist)
		s = s.Where("photos.photo_lat BETWEEN ? AND ?", latMin, latMax)
	}

	if f.Lng > 0 {
		lngMin := f.Lng - SearchRadius*float32(f.Dist)
		lngMax := f.Lng + SearchRadius*float32(f.Dist)
		s = s.Where("photos.photo_lng BETWEEN ? AND ?", lngMin, lngMax)
	}

	if !f.Before.IsZero() {
		s = s.Where("photos.taken_at <= ?", f.Before.Format("2006-01-02"))
	}

	if !f.After.IsZero() {
		s = s.Where("photos.taken_at >= ?", f.After.Format("2006-01-02"))
	}

	switch f.Order {
	case entity.SortOrderRelevance:
		if f.Label != "" {
			s = s.Order("photo_quality DESC, photos_labels.uncertainty ASC, taken_at DESC, files.file_primary DESC")
		} else {
			s = s.Order("photo_quality DESC, taken_at DESC, files.file_primary DESC")
		}
	case entity.SortOrderNewest:
		s = s.Order("taken_at DESC, photos.photo_uuid, files.file_primary DESC")
	case entity.SortOrderOldest:
		s = s.Order("taken_at, photos.photo_uuid, files.file_primary DESC")
	case entity.SortOrderImported:
		s = s.Order("photos.id DESC, files.file_primary DESC")
	case entity.SortOrderSimilar:
		s = s.Order("files.file_main_color, photos.location_id, files.file_diff, taken_at DESC, files.file_primary DESC")
	default:
		s = s.Order("taken_at DESC, photos.photo_uuid, files.file_primary DESC")
	}

	if f.Count > 0 && f.Count <= 1000 {
		s = s.Limit(f.Count).Offset(f.Offset)
	} else {
		s = s.Limit(100).Offset(0)
	}

	if result := s.Scan(&results); result.Error != nil {
		return results, 0, result.Error
	}

	if f.Merged {
		return results.Merged()
	}

	return results, len(results), nil
}

// PhotoByID returns a Photo based on the ID.
func (q *Query) PhotoByID(photoID uint64) (photo entity.Photo, err error) {
	if err := q.db.Unscoped().Where("id = ?", photoID).
		Preload("Links").
		Preload("Description").
		Preload("Location").
		Preload("Location.Place").
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		First(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

// PhotoByUUID returns a Photo based on the UUID.
func (q *Query) PhotoByUUID(photoUUID string) (photo entity.Photo, err error) {
	if err := q.db.Unscoped().Where("photo_uuid = ?", photoUUID).
		Preload("Links").
		Preload("Description").
		Preload("Location").
		Preload("Location.Place").
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		First(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

// PreloadPhotoByUUID returns a Photo based on the UUID with all dependencies preloaded.
func (q *Query) PreloadPhotoByUUID(photoUUID string) (photo entity.Photo, err error) {
	if err := q.db.Unscoped().Where("photo_uuid = ?", photoUUID).
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		Preload("Camera").
		Preload("Lens").
		Preload("Links").
		Preload("Location").
		Preload("Location.Place").
		Preload("Description").
		First(&photo).Error; err != nil {
		return photo, err
	}

	photo.PreloadMany()

	return photo, nil
}

// MissingPhotos returns photo entities without existing files.
func (q *Query) MissingPhotos(limit int, offset int) (entities []entity.Photo, err error) {
	err = q.db.
		Select("photos.*").
		Joins("JOIN files a ON photos.id = a.photo_id ").
		Joins("LEFT JOIN files b ON a.photo_id = b.photo_id AND a.id != b.id AND b.file_missing = 0").
		Where("a.file_missing = 1 AND b.id IS NULL").
		Group("photos.id").
		Limit(limit).Offset(offset).Find(&entities).Error

	return entities, err
}

// ResetPhotosQuality resets the quality of photos without primary file to -1.
func (q *Query) ResetPhotosQuality() error {
	return q.db.Table("photos").
		Where("id IN (SELECT photos.id FROM photos LEFT JOIN files ON photos.id = files.photo_id AND files.file_primary = 1 WHERE files.id IS NULL GROUP BY photos.id)").
		Update("photo_quality", -1).Error
}
