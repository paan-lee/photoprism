package query

import (
	"errors"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/form"
)

// PhotoSelection returns all selected photos.
func (q *Query) PhotoSelection(f form.Selection) (results []entity.Photo, err error) {
	if f.Empty() {
		return results, errors.New("no photos selected")
	}

	s := q.db.NewScope(nil).DB()

	s = s.Table("photos").
		Select("photos.*").
		Joins("LEFT JOIN photos_labels ON photos_labels.photo_id = photos.id").
		Joins("LEFT JOIN labels ON photos_labels.label_id = labels.id AND labels.deleted_at IS NULL").
		Where("photos.deleted_at IS NULL").
		Group("photos.id")

	s = s.Where("photos.photo_uuid IN (?) OR labels.label_uuid IN (?)", f.Photos, f.Labels)

	if result := s.Scan(&results); result.Error != nil {
		return results, result.Error
	}

	return results, nil
}
