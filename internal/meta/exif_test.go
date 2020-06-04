package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExif(t *testing.T) {
	t.Run("photoshop.jpg", func(t *testing.T) {
		data, err := Exif("testdata/photoshop.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Michael Mayer", data.Artist)
		assert.Equal(t, "2020-01-01T16:28:23Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-01-01T17:28:23Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "Example file for development", data.Description)
		assert.Equal(t, "This is a legal notice", data.Copyright)
		assert.Equal(t, 540, data.Height)
		assert.Equal(t, 720, data.Width)
		assert.Equal(t, float32(52.45969), data.Lat)
		assert.Equal(t, float32(13.321832), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/50", data.Exposure)
		assert.Equal(t, "HUAWEI", data.CameraMake)
		assert.Equal(t, "ELE-L29", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 27, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))

		// TODO: Values are empty - why?
		// assert.Equal(t, "HUAWEI P30 Rear Main Camera", data.LensModel)
	})

	t.Run("ladybug.jpg", func(t *testing.T) {
		data, err := Exif("testdata/ladybug.jpg")

		if err != nil {
			t.Fatal(err)
		}

		//  t.Logf("all: %+v", data.All)

		assert.Equal(t, "Photographer: TMB", data.Artist)
		assert.Equal(t, "2011-07-10T17:34:28Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2011-07-10T19:34:28Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)    // Should be "Ladybug"
		assert.Equal(t, "", data.Keywords) // Should be "Ladybug"
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 540, data.Height)
		assert.Equal(t, 720, data.Width)
		assert.Equal(t, float32(51.254852), data.Lat)
		assert.Equal(t, float32(7.389468), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/125", data.Exposure)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS 50D", data.CameraModel)
		assert.Equal(t, "Thomas Meyer-Boudnik", data.CameraOwner)
		assert.Equal(t, "2260716910", data.CameraSerial)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "EF100mm f/2.8 Macro USM", data.LensModel)
		assert.Equal(t, 100, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("gopro_hd2.jpg", func(t *testing.T) {
		data, err := Exif("testdata/gopro_hd2.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2017-12-21T05:17:28Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2017-12-21T05:17:28Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords)
		assert.Equal(t, "DCIM\\100GOPRO", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 180, data.Height)
		assert.Equal(t, 240, data.Width)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/2462", data.Exposure)
		assert.Equal(t, "GoPro", data.CameraMake)
		assert.Equal(t, "HD2", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 16, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("tweethog.png", func(t *testing.T) {
		_, err := Exif("testdata/tweethog.png")

		if err == nil {
			t.Fatal("err should NOT be nil")
		}

		assert.Equal(t, "no exif data in tweethog.png", err.Error())
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		data, err := Exif("testdata/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2018-09-10T03:16:13Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2018-09-10T12:16:13Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, float32(34.79745), data.Lat)
		assert.Equal(t, float32(134.76463), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/4000", data.Exposure)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 7", data.CameraModel)
		assert.Equal(t, 74, data.FocalLength)
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, "Apple", data.LensMake)
		assert.Equal(t, "iPhone 7 back camera 3.99mm f/1.8", data.LensModel)

	})

	t.Run("gps-2000.jpg", func(t *testing.T) {
		data, err := Exif("testdata/gps-2000.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("GPS 2000: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.True(t, data.TakenAt.IsZero())
		assert.True(t, data.TakenAtLocal.IsZero())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 0, data.Height) // TODO
		assert.Equal(t, 0, data.Width)  // TODO
		assert.Equal(t, float32(-38.405193), data.Lat)
		assert.Equal(t, float32(144.18896), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "", data.Exposure)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("image-2011.jpg", func(t *testing.T) {
		data, err := Exif("testdata/image-2011.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("ALL: %+v", data.All)

		/*
		  Exiftool date information:

		  File Modification Date/Time     : 2020:05:15 08:25:46+00:00
		  File Access Date/Time           : 2020:05:15 08:25:47+00:00
		  File Inode Change Date/Time     : 2020:05:15 08:25:46+00:00
		  Modify Date                     : 2020:05:15 10:25:45
		  Create Date                     : 2011:07:19 11:36:38
		  Metadata Date                   : 2020:05:15 10:25:45+02:00

		*/

		// assert.Equal(t, "2011-07-19T11:36:38Z", data.TakenAt.Format("2006-01-02T15:04:05Z")) // TODO
		// assert.Equal(t, "2011-07-19T11:36:38Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))  // TODO
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/1100", data.Exposure)
		assert.Equal(t, "SAMSUNG", data.CameraMake)
		assert.Equal(t, "GT-I9000", data.CameraModel)
		assert.Equal(t, 3, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("ship.jpg", func(t *testing.T) {
		data, err := Exif("testdata/ship.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2019-05-12T15:13:53Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2019-05-12T17:13:53Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, float32(53.12349), data.Lat)
		assert.Equal(t, float32(18.00152), data.Lng)
		assert.Equal(t, 63, data.Altitude)
		assert.Equal(t, "1/100", data.Exposure)
		assert.Equal(t, "Xiaomi", data.CameraMake)
		assert.Equal(t, "Mi A1", data.CameraModel)
		assert.Equal(t, 52, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("no-exif-data.jpg", func(t *testing.T) {
		_, err := Exif("testdata/no-exif-data.jpg")

		if err == nil {
			t.Fatal("err should NOT be nil")
		}

		assert.Equal(t, "no exif data in no-exif-data.jpg", err.Error())
	})

	t.Run("screenshot.png", func(t *testing.T) {
		data, err := Exif("testdata/screenshot.png")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "721", data.All["PixelXDimension"])
		assert.Equal(t, "332", data.All["PixelYDimension"])
	})

	t.Run("orientation.jpg", func(t *testing.T) {
		data, err := Exif("testdata/orientation.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "3264", data.All["PixelXDimension"])
		assert.Equal(t, "1836", data.All["PixelYDimension"])
		assert.Equal(t, 3264, data.Width)
		assert.Equal(t, 1836, data.Height)
		assert.Equal(t, 6, data.Orientation) // TODO: Should be 1

		if err := data.JSON("testdata/orientation.json", "orientation.jpg"); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 326, data.Width)
		assert.Equal(t, 184, data.Height)
		assert.Equal(t, 1, data.Orientation)

		if err := data.JSON("testdata/orientation.json", "foo.jpg"); err != nil {
			assert.EqualError(t, err, "meta: original name foo.jpg does not match orientation.jpg (json)")
		} else {
			t.Error("error expected when providing wrong orginal name")
		}
	})

	t.Run("gopher-preview.jpg", func(t *testing.T) {
		_, err := Exif("testdata/gopher-preview.jpg")

		assert.EqualError(t, err, "no exif data in gopher-preview.jpg")
	})
}
