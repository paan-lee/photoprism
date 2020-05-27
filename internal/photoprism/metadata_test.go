package photoprism

import (
	"os"
	"testing"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/meta"
	"github.com/stretchr/testify/assert"
)

func TestMediaFile_Exif_JPEG(t *testing.T) {
	conf := config.TestConfig()

	t.Run("elephants.jpg", func(t *testing.T) {
		img, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		assert.Nil(t, err)

		data := img.MetaData()

		assert.Empty(t, err)

		assert.IsType(t, meta.Data{}, data)

		assert.Equal(t, "", data.DocumentID)
		assert.Equal(t, "2013-11-26 13:53:55 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "2013-11-26 15:53:55 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "Canon EOS 6D", data.CameraModel)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "EF70-200mm f/4L IS USM", data.LensModel)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "Africa/Johannesburg", data.TimeZone)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, 111, data.FocalLength)
		assert.Equal(t, "1/640", data.Exposure)
		assert.Equal(t, float32(6.644), data.Aperture)
		assert.Equal(t, float32(10), data.FNumber)
		assert.Equal(t, 200, data.Iso)
		assert.Equal(t, float32(-33.45347), data.Lat)
		assert.Equal(t, float32(25.764645), data.Lng)
		assert.Equal(t, 190, data.Altitude)
		assert.Equal(t, 497, data.Width)
		assert.Equal(t, 331, data.Height)
		assert.Equal(t, false, data.Flash)
		assert.Equal(t, "", data.Description)
		t.Logf("UTC: %s", data.TakenAt.String())
		t.Logf("Local: %s", data.TakenAtLocal.String())
	})

	t.Run("fern_green.jpg", func(t *testing.T) {
		img, err := NewMediaFile(conf.ExamplesPath() + "/fern_green.jpg")

		assert.Nil(t, err)

		info := img.MetaData()

		assert.Empty(t, err)

		assert.IsType(t, meta.Data{}, info)

		assert.Equal(t, "", info.DocumentID)
		assert.Equal(t, 1, info.Orientation)
		assert.Equal(t, "Canon EOS 7D", info.CameraModel)
		assert.Equal(t, "Canon", info.CameraMake)
		assert.Equal(t, "EF100mm f/2.8L Macro IS USM", info.LensModel)
		assert.Equal(t, "", info.LensMake)
		assert.Equal(t, "", info.TimeZone)
		assert.Equal(t, "", info.Artist)
		assert.Equal(t, 100, info.FocalLength)
		assert.Equal(t, "1/250", info.Exposure)
		assert.Equal(t, float32(6.644), info.Aperture)
		assert.Equal(t, float32(10), info.FNumber)
		assert.Equal(t, 200, info.Iso)
		assert.Equal(t, 0, info.Altitude)
		assert.Equal(t, 331, info.Width)
		assert.Equal(t, 331, info.Height)
		assert.Equal(t, true, info.Flash)
		assert.Equal(t, "", info.Description)
		t.Logf("UTC: %s", info.TakenAt.String())
		t.Logf("Local: %s", info.TakenAtLocal.String())
	})
}

func TestMediaFile_Exif_DNG(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	img, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")

	assert.Nil(t, err)

	info := img.MetaData()

	assert.Empty(t, err)

	assert.IsType(t, meta.Data{}, info)

	assert.Equal(t, "", info.DocumentID)
	assert.Equal(t, "2019-06-06 07:29:51 +0000 UTC", info.TakenAt.String())
	assert.Equal(t, "2019-06-06 07:29:51 +0000 UTC", info.TakenAtLocal.String())
	assert.Equal(t, 1, info.Orientation)
	assert.Equal(t, "Canon EOS 6D", info.CameraModel)
	assert.Equal(t, "Canon", info.CameraMake)
	assert.Equal(t, "EF24-105mm f/4L IS USM", info.LensModel)
	assert.Equal(t, "", info.Artist)
	assert.Equal(t, 65, info.FocalLength)
	assert.Equal(t, "1/60", info.Exposure)
	assert.Equal(t, float32(4.971), info.Aperture)
	assert.Equal(t, 1000, info.Iso)
	assert.Equal(t, float32(0), info.Lat)
	assert.Equal(t, float32(0), info.Lng)
	assert.Equal(t, 0, info.Altitude)
	assert.Equal(t, 256, info.Width)
	assert.Equal(t, 171, info.Height)
	assert.Equal(t, false, info.Flash)
	assert.Equal(t, "", info.Description)
}

func TestMediaFile_Exif_HEIF(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	img, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

	assert.Nil(t, err)

	info := img.MetaData()

	assert.IsType(t, meta.Data{}, info)

	assert.Nil(t, err)

	convert := NewConvert(conf)

	jpeg, err := convert.ToJpeg(img, true)

	assert.Nil(t, err)

	jpegInfo := jpeg.MetaData()

	assert.IsType(t, meta.Data{}, jpegInfo)

	assert.Nil(t, err)

	assert.Equal(t, "", jpegInfo.DocumentID)
	assert.Equal(t, "2018-09-10 03:16:13 +0000 UTC", jpegInfo.TakenAt.String())
	assert.Equal(t, "2018-09-10 12:16:13 +0000 UTC", jpegInfo.TakenAtLocal.String())
	assert.Equal(t, 6, jpegInfo.Orientation)
	assert.Equal(t, "iPhone 7", jpegInfo.CameraModel)
	assert.Equal(t, "Apple", jpegInfo.CameraMake)
	assert.Equal(t, "iPhone 7 back camera 3.99mm f/1.8", jpegInfo.LensModel)
	assert.Equal(t, "Apple", jpegInfo.LensMake)
	assert.Equal(t, "Asia/Tokyo", jpegInfo.TimeZone)
	assert.Equal(t, "", jpegInfo.Artist)
	assert.Equal(t, 74, jpegInfo.FocalLength)
	assert.Equal(t, "1/4000", jpegInfo.Exposure)
	assert.Equal(t, float32(1.696), jpegInfo.Aperture)
	assert.Equal(t, 20, jpegInfo.Iso)
	assert.Equal(t, float32(34.79745), jpegInfo.Lat)
	assert.Equal(t, float32(134.76463), jpegInfo.Lng)
	assert.Equal(t, 0, jpegInfo.Altitude)
	assert.Equal(t, 4032, jpegInfo.Width)
	assert.Equal(t, 3024, jpegInfo.Height)
	assert.Equal(t, false, jpegInfo.Flash)
	assert.Equal(t, "", jpegInfo.Description)

	if err := os.Remove(conf.ExamplesPath() + "/.photoprism/iphone_7.jpg"); err != nil {
		t.Error(err)
	}
}
