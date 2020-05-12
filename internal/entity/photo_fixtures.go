package entity

import (
	"time"
)

var editTime = time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)

type PhotoMap map[string]Photo

func (m PhotoMap) Get(name string) Photo {
	if result, ok := m[name]; ok {
		return result
	}

	return Photo{PhotoName: name}
}

func (m PhotoMap) Pointer(name string) *Photo {
	if result, ok := m[name]; ok {
		return &result
	}

	return &Photo{PhotoName: name}
}

var PhotoFixtures = PhotoMap{
	"19800101_000002_D640C559": {
		ID:               1000000,
		PhotoUUID:        "pt9jtdre2lvl0yh7",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "2790/02",
		PhotoName:        "19800101_000002_D640C559",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2790,
		PhotoMonth:       2,
		Description:      DescriptionFixtures.Get("lake", 1000000),
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums: []Album{
			AlbumFixtures.Get("holiday-2030"),
		},
		Files: []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(1000000, "flower", 38, "image"),
			LabelFixtures.PhotoLabel(1000000, "cake", 38, "manual"),
		},
		CreatedAt: time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo01": {
		ID:               1000001,
		PhotoUUID:        "pt9jtdre2lvl0yh8",
		TakenAt:          time.Date(2006, 1, 1, 2, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2006, 1, 1, 2, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "2790/02",
		PhotoName:        "Photo01",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    true,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2790,
		PhotoMonth:       2,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           CameraFixtures.Pointer("canon-eos-6d"),
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo02": {
		ID:               1000002,
		PhotoUUID:        "pt9jtdre2lvl0yh9",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/03",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        1990,
		PhotoMonth:       3,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000002, "cake", 20, "image")},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo03": {
		ID:               1000003,
		PhotoUUID:        "pt9jtdre2lvl0yh0",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Time{},
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/04",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        1990,
		PhotoMonth:       4,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels: []PhotoLabel{
			LabelFixtures.PhotoLabel(1000003, "cow", 20, "image"),
			LabelFixtures.PhotoLabel(1000003, "updatePhotoLabel", 20, "manual"),
		},
		CreatedAt: time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo04": {
		ID:               1000004,
		PhotoUUID:        "pt9jtdre2lvl0y11",
		TakenAt:          time.Date(2014, 7, 17, 15, 42, 12, 0, time.UTC),
		TakenAtLocal:     time.Date(2014, 7, 17, 15, 42, 12, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "Neckarbrücke",
		TitleSrc:         "",
		PhotoPath:        "2014/07",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords: []Keyword{
			KeywordFixtures.Get("bridge"),
		},
		Albums: []Album{
			AlbumFixtures.Get("berlin-2019"),
		},
		Files:     []File{},
		Labels:    []PhotoLabel{LabelFixtures.PhotoLabel(1000004, "batchdelete", 20, "image")},
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:  nil,
		DeletedAt: nil,
	},
	"Photo05": {
		ID:               1000005,
		PhotoUUID:        "pt9jtdre2lvl0y12",
		TakenAt:          time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "Reunion",
		TitleSrc:         "",
		PhotoPath:        "2014/07",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "123",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000005, "updateLabel", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo06": {
		ID:               1000006,
		PhotoUUID:        "pt9jtdre2lvl0y13",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "ToBeUpdated",
		TitleSrc:         "exif",
		PhotoPath:        "2016/11",
		PhotoName:        "UpdatePhoto",
		PhotoQuality:     0,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000006, "updatePhotoLabel", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo07": {
		ID:               1000007,
		PhotoUUID:        "pt9jtdre2lvl0y14",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "ToBeUpdated",
		TitleSrc:         "exif",
		PhotoPath:        "2016/11",
		PhotoName:        "UpdatePhoto",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000007, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         &editTime,
		DeletedAt:        nil,
	},
	"Photo08": {
		ID:               1000008,
		PhotoUUID:        "pt9jtdre2lvl0y15",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Black beach",
		TitleSrc:         "exif",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "85d1ea7d382c",
		LocationID:       "85d1ea7d382c",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo09": {
		ID:               1000009,
		PhotoUUID:        "pt9jtdre2lvl0y16",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "85d1ea7d382c",
		LocationID:       "85d1ea7d382c",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         &LocationFixturesMexico,
		Place:            PlaceFixtures.Pointer("teotihuacan"),
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo10": {
		ID:               1000010,
		PhotoUUID:        "pt9jtdre2lvl0y17",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         &LocationFixturesHassloch,
		Place:            PlaceFixtures.Pointer("holidaypark"),
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo11": {
		ID:               1000011,
		PhotoUUID:        "pt9jtdre2lvl0y18",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         &LocationFixturesEmptyNameLongCity,
		Place:            PlaceFixtures.Pointer("emptyNameLongCity"),
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo12": {
		ID:               1000012,
		PhotoUUID:        "pt9jtdre2lvl0y19",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         &LocationFixturesEmptyNameShortCity,
		Place:            PlaceFixtures.Pointer("emptyNameShortCity"),
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo13": {
		ID:               1000013,
		PhotoUUID:        "pt9jtdre2lvl0y20",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         &LocationFixturesVeryLongLocName,
		Place:            PlaceFixtures.Pointer("veryLongLocName"),
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000008, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo14": {
		ID:               1000014,
		PhotoUUID:        "pt9jtdre2lvl0y21",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "Title",
		TitleSrc:         "",
		PhotoPath:        "2016/11",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         &LocationFixturesMediumLongLocName,
		Place:            PlaceFixtures.Pointer("mediumLongLocName"),
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(1000014, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo15": {
		ID:               1000015,
		PhotoUUID:        "pt9jtdre2lvl0y22",
		TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "location",
		PhotoTitle:       "TitleToBeSet",
		TitleSrc:         "location",
		PhotoPath:        "",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         1.234,
		PhotoLng:         4.321,
		PhotoAltitude:    3,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "location",
		TimeZone:         "",
		PhotoCountry:     "",
		PhotoYear:        0,
		PhotoMonth:       0,
		Description:      DescriptionFixtures.Get("lake", 1000015),
		DescriptionSrc:   "location",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(10000015, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo16": {
		ID:               1000016,
		PhotoUUID:        "pt9jtdre2lvl0y23",
		TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "ForDeletion",
		TitleSrc:         "",
		PhotoPath:        "",
		PhotoName:        "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoVideo:       false,
		PhotoLat:         1.234,
		PhotoLng:         4.321,
		PhotoAltitude:    3,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "",
		LocationID:       "",
		LocationSrc:      "location",
		TimeZone:         "",
		PhotoCountry:     "",
		PhotoYear:        0,
		PhotoMonth:       0,
		Description:      DescriptionFixtures.Get("lake", 1000015),
		DescriptionSrc:   "location",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{LabelFixtures.PhotoLabel(10000015, "landscape", 20, "image")},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
}

// CreatePhotoFixtures inserts known entities into the database for testing.
func CreatePhotoFixtures() {
	for _, entity := range PhotoFixtures {
		Db().Create(&entity)
	}
}
