package query

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPhotosResults_Merged(t *testing.T) {
	result1 := PhotoResult{
		ID:               111111,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        time.Time{},
		TakenAt:          time.Time{},
		TakenAtLocal:     time.Time{},
		TakenSrc:         "",
		TimeZone:         "",
		PhotoUID:         "",
		PhotoPath:        "",
		PhotoName:        "",
		PhotoTitle:       "Photo1",
		PhotoYear:        0,
		PhotoMonth:       0,
		PhotoCountry:     "",
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		Merged:           false,
		CameraID:         0,
		CameraModel:      "",
		CameraMake:       "",
		LensID:           0,
		LensModel:        "",
		LensMake:         "",
		CellID:           "",
		PlaceID:          "",
		PlaceLabel:       "",
		PlaceCity:        "",
		PlaceState:       "",
		PlaceCountry:     "",
		FileID:           0,
		FileUID:          "",
		FilePrimary:      false,
		FileMissing:      false,
		FileName:         "",
		FileHash:         "",
		FileType:         "",
		FileMime:         "",
		FileWidth:        0,
		FileHeight:       0,
		FileOrientation:  0,
		FileAspectRatio:  0,
		FileColors:       "",
		FileChroma:       0,
		FileLuminance:    "",
		FileDiff:         0,
		Files:            nil,
	}

	result2 := PhotoResult{
		ID:               22222,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        time.Time{},
		TakenAt:          time.Time{},
		TakenAtLocal:     time.Time{},
		TakenSrc:         "",
		TimeZone:         "",
		PhotoUID:         "",
		PhotoPath:        "",
		PhotoName:        "",
		PhotoTitle:       "Photo2",
		PhotoYear:        0,
		PhotoMonth:       0,
		PhotoCountry:     "",
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		Merged:           false,
		CameraID:         0,
		CameraModel:      "",
		CameraMake:       "",
		LensID:           0,
		LensModel:        "",
		LensMake:         "",
		CellID:           "",
		PlaceID:          "",
		PlaceLabel:       "",
		PlaceCity:        "",
		PlaceState:       "",
		PlaceCountry:     "",
		FileID:           0,
		FileUID:          "",
		FilePrimary:      false,
		FileMissing:      false,
		FileName:         "",
		FileHash:         "",
		FileType:         "",
		FileMime:         "",
		FileWidth:        0,
		FileHeight:       0,
		FileOrientation:  0,
		FileAspectRatio:  0,
		FileColors:       "",
		FileChroma:       0,
		FileLuminance:    "",
		FileDiff:         0,
		Files:            nil,
	}

	results := PhotoResults{result1, result2}

	merged, count, err := results.Merged()

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, count)
	t.Log(merged)
}
func TestPhotosResults_UIDs(t *testing.T) {
	result1 := PhotoResult{
		ID:               111111,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        time.Time{},
		TakenAt:          time.Time{},
		TakenAtLocal:     time.Time{},
		TakenSrc:         "",
		TimeZone:         "",
		PhotoUID:         "123",
		PhotoPath:        "",
		PhotoName:        "",
		PhotoTitle:       "Photo1",
		PhotoYear:        0,
		PhotoMonth:       0,
		PhotoCountry:     "",
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		Merged:           false,
		CameraID:         0,
		CameraModel:      "",
		CameraMake:       "",
		LensID:           0,
		LensModel:        "",
		LensMake:         "",
		CellID:           "",
		PlaceID:          "",
		PlaceLabel:       "",
		PlaceCity:        "",
		PlaceState:       "",
		PlaceCountry:     "",
		FileID:           0,
		FileUID:          "",
		FilePrimary:      false,
		FileMissing:      false,
		FileName:         "",
		FileHash:         "",
		FileType:         "",
		FileMime:         "",
		FileWidth:        0,
		FileHeight:       0,
		FileOrientation:  0,
		FileAspectRatio:  0,
		FileColors:       "",
		FileChroma:       0,
		FileLuminance:    "",
		FileDiff:         0,
		Files:            nil,
	}

	result2 := PhotoResult{
		ID:               22222,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        time.Time{},
		TakenAt:          time.Time{},
		TakenAtLocal:     time.Time{},
		TakenSrc:         "",
		TimeZone:         "",
		PhotoUID:         "456",
		PhotoPath:        "",
		PhotoName:        "",
		PhotoTitle:       "Photo2",
		PhotoYear:        0,
		PhotoMonth:       0,
		PhotoCountry:     "",
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoLat:         0,
		PhotoLng:         0,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		PhotoQuality:     0,
		PhotoResolution:  0,
		Merged:           false,
		CameraID:         0,
		CameraModel:      "",
		CameraMake:       "",
		LensID:           0,
		LensModel:        "",
		LensMake:         "",
		CellID:           "",
		PlaceID:          "",
		PlaceLabel:       "",
		PlaceCity:        "",
		PlaceState:       "",
		PlaceCountry:     "",
		FileID:           0,
		FileUID:          "",
		FilePrimary:      false,
		FileMissing:      false,
		FileName:         "",
		FileHash:         "",
		FileType:         "",
		FileMime:         "",
		FileWidth:        0,
		FileHeight:       0,
		FileOrientation:  0,
		FileAspectRatio:  0,
		FileColors:       "",
		FileChroma:       0,
		FileLuminance:    "",
		FileDiff:         0,
		Files:            nil,
	}

	results := PhotoResults{result1, result2}

	result := results.UIDs()
	assert.Equal(t, []string{"123", "456"}, result)
}

func TestPhotosResult_ShareFileName(t *testing.T) {
	t.Run("with photo title", func(t *testing.T) {
		result1 := PhotoResult{
			ID:               111111,
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			DeletedAt:        time.Time{},
			TakenAt:          time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
			TakenAtLocal:     time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
			TakenSrc:         "",
			TimeZone:         "",
			PhotoUID:         "uid123",
			PhotoPath:        "",
			PhotoName:        "",
			PhotoTitle:       "PhotoTitle123",
			PhotoYear:        0,
			PhotoMonth:       0,
			PhotoCountry:     "",
			PhotoFavorite:    false,
			PhotoPrivate:     false,
			PhotoLat:         0,
			PhotoLng:         0,
			PhotoAltitude:    0,
			PhotoIso:         0,
			PhotoFocalLength: 0,
			PhotoFNumber:     0,
			PhotoExposure:    "",
			PhotoQuality:     0,
			PhotoResolution:  0,
			Merged:           false,
			CameraID:         0,
			CameraModel:      "",
			CameraMake:       "",
			LensID:           0,
			LensModel:        "",
			LensMake:         "",
			CellID:           "",
			PlaceID:          "",
			PlaceLabel:       "",
			PlaceCity:        "",
			PlaceState:       "",
			PlaceCountry:     "",
			FileID:           0,
			FileUID:          "",
			FilePrimary:      false,
			FileMissing:      false,
			FileName:         "",
			FileHash:         "",
			FileType:         "",
			FileMime:         "",
			FileWidth:        0,
			FileHeight:       0,
			FileOrientation:  0,
			FileAspectRatio:  0,
			FileColors:       "",
			FileChroma:       0,
			FileLuminance:    "",
			FileDiff:         0,
			Files:            nil,
		}

		r := result1.ShareBase(0)
		assert.Contains(t, r, "20131111-090718-Phototitle123")
	})
	t.Run("without photo title", func(t *testing.T) {
		result1 := PhotoResult{
			ID:               111111,
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			DeletedAt:        time.Time{},
			TakenAt:          time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC),
			TakenAtLocal:     time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
			TakenSrc:         "",
			TimeZone:         "",
			PhotoUID:         "uid123",
			PhotoPath:        "",
			PhotoName:        "",
			PhotoTitle:       "",
			PhotoYear:        0,
			PhotoMonth:       0,
			PhotoCountry:     "",
			PhotoFavorite:    false,
			PhotoPrivate:     false,
			PhotoLat:         0,
			PhotoLng:         0,
			PhotoAltitude:    0,
			PhotoIso:         0,
			PhotoFocalLength: 0,
			PhotoFNumber:     0,
			PhotoExposure:    "",
			PhotoQuality:     0,
			PhotoResolution:  0,
			Merged:           false,
			CameraID:         0,
			CameraModel:      "",
			CameraMake:       "",
			LensID:           0,
			LensModel:        "",
			LensMake:         "",
			CellID:           "",
			PlaceID:          "",
			PlaceLabel:       "",
			PlaceCity:        "",
			PlaceState:       "",
			PlaceCountry:     "",
			FileID:           0,
			FileUID:          "",
			FilePrimary:      false,
			FileMissing:      false,
			FileName:         "",
			FileHash:         "",
			FileType:         "",
			FileMime:         "",
			FileWidth:        0,
			FileHeight:       0,
			FileOrientation:  0,
			FileAspectRatio:  0,
			FileColors:       "",
			FileChroma:       0,
			FileLuminance:    "",
			FileDiff:         0,
			Files:            nil,
		}

		r := result1.ShareBase(0)
		assert.Contains(t, r, "20151111-090718-uid123")
	})
}
