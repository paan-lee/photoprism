package form

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPhoto(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		photo := Photo{
			TakenAt:          time.Date(2008, 1, 1, 2, 0, 0, 0, time.UTC),
			TakenAtLocal:     time.Date(2008, 1, 1, 2, 0, 0, 0, time.UTC),
			TakenSrc:         "meta",
			TimeZone:         "UTC",
			PhotoTitle:       "Black beach",
			TitleSrc:         "manual",
			PhotoFavorite:    false,
			PhotoPrivate:     false,
			PhotoType:        "image",
			PhotoReview:      false,
			PhotoLat:         9.9999,
			PhotoLng:         8.8888,
			PhotoAltitude:    2,
			PhotoIso:         5,
			PhotoFocalLength: 10,
			PhotoFNumber:     3.3,
			PhotoExposure:    "exposure",
			CameraID:         uint(3),
			CameraSrc:        "meta",
			LensID:           uint(6),
			GeoID:            "1234",
			GeoSrc:           "geo",
			PlaceID:          "765",
			PhotoCountry:     "de"}

		r, err := NewPhoto(photo)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, time.Date(2008, 1, 1, 2, 0, 0, 0, time.UTC), r.TakenAt)
		assert.Equal(t, time.Date(2008, 1, 1, 2, 0, 0, 0, time.UTC), r.TakenAtLocal)
		assert.Equal(t, "meta", r.TakenSrc)
		assert.Equal(t, "UTC", r.TimeZone)
		assert.Equal(t, "Black beach", r.PhotoTitle)
		assert.Equal(t, "manual", r.TitleSrc)
		assert.Equal(t, false, r.PhotoFavorite)
		assert.Equal(t, false, r.PhotoPrivate)
		assert.Equal(t, "image", r.PhotoType)
		assert.Equal(t, false, r.PhotoReview)
		assert.Equal(t, float32(9.9999), r.PhotoLat)
		assert.Equal(t, float32(8.8888), r.PhotoLng)
		assert.Equal(t, 2, r.PhotoAltitude)
		assert.Equal(t, 5, r.PhotoIso)
		assert.Equal(t, 10, r.PhotoFocalLength)
		assert.Equal(t, float32(3.3), r.PhotoFNumber)
		assert.Equal(t, "exposure", r.PhotoExposure)
		assert.Equal(t, uint(3), r.CameraID)
		assert.Equal(t, "meta", r.CameraSrc)
		assert.Equal(t, uint(6), r.LensID)
		assert.Equal(t, "1234", r.GeoID)
		assert.Equal(t, "geo", r.GeoSrc)
		assert.Equal(t, "765", r.PlaceID)
		assert.Equal(t, "de", r.PhotoCountry)
	})
}
