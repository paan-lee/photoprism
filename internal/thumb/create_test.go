package thumb

import (
	"os"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/stretchr/testify/assert"
)

func TestResampleOptions(t *testing.T) {
	t.Run("ResamplePng, FillCenter", func(t *testing.T) {
		method, filter, format := ResampleOptions(ResamplePng, ResampleFillCenter, ResampleDefault)

		assert.Equal(t, ResampleFillCenter, method)
		assert.Equal(t, imaging.Lanczos.Support, filter.Support)
		assert.Equal(t, fs.FormatPng, format)
	})
	t.Run("ResampleNearestNeighbor, FillTopLeft", func(t *testing.T) {
		method, filter, format := ResampleOptions(ResampleNearestNeighbor, ResampleFillTopLeft)

		assert.Equal(t, ResampleFillTopLeft, method)
		assert.Equal(t, imaging.NearestNeighbor.Support, filter.Support)
		assert.Equal(t, fs.FormatJpeg, format)
	})
	t.Run("ResampleNearestNeighbor, FillBottomRight", func(t *testing.T) {
		method, filter, format := ResampleOptions(ResampleNearestNeighbor, ResampleFillBottomRight)

		assert.Equal(t, ResampleFillBottomRight, method)
		assert.Equal(t, imaging.NearestNeighbor.Support, filter.Support)
		assert.Equal(t, fs.FormatJpeg, format)
	})
}

func TestResample(t *testing.T) {
	t.Run("tile50 options", func(t *testing.T) {
		tile50 := Types["tile_50"]

		src := "testdata/example.jpg"

		assert.FileExists(t, src)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		result := Resample(img, tile50.Width, tile50.Height, tile50.Options...)

		boundsNew := result.Bounds()

		assert.Equal(t, 50, boundsNew.Max.X)
		assert.Equal(t, 50, boundsNew.Max.Y)
	})
	t.Run("left_224 options", func(t *testing.T) {
		left224 := Types["left_224"]

		src := "testdata/example.jpg"

		assert.FileExists(t, src)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		result := Resample(img, left224.Width, left224.Height, left224.Options...)

		boundsNew := result.Bounds()

		assert.Equal(t, 224, boundsNew.Max.X)
		assert.Equal(t, 224, boundsNew.Max.Y)
	})
	t.Run("right_224 options", func(t *testing.T) {
		right224 := Types["right_224"]

		src := "testdata/example.jpg"

		assert.FileExists(t, src)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		result := Resample(img, right224.Width, right224.Height, right224.Options...)

		boundsNew := result.Bounds()

		assert.Equal(t, 224, boundsNew.Max.X)
		assert.Equal(t, 224, boundsNew.Max.Y)
	})
	t.Run("fit_1280 options", func(t *testing.T) {
		fit1280 := Types["fit_1280"]

		src := "testdata/example.jpg"

		assert.FileExists(t, src)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		result := Resample(img, fit1280.Width, fit1280.Height, fit1280.Options...)

		boundsNew := result.Bounds()

		assert.Equal(t, 750, boundsNew.Max.X)
		assert.Equal(t, 500, boundsNew.Max.Y)
	})
}

func TestPostfix(t *testing.T) {
	tile50 := Types["tile_50"]

	result := Postfix(tile50.Width, tile50.Height, tile50.Options...)

	assert.Equal(t, "50x50_center.jpg", result)
}

func TestFilename(t *testing.T) {
	t.Run("colors", func(t *testing.T) {
		colorThumb := Types["colors"]

		result, err := Filename("123456789098765432", "testdata", colorThumb.Width, colorThumb.Height, colorThumb.Options...)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "testdata/1/2/3/123456789098765432_3x3_resize.png", result)
	})

	t.Run("fit_720", func(t *testing.T) {
		fit720 := Types["fit_720"]

		result, err := Filename("123456789098765432", "testdata", fit720.Width, fit720.Height, fit720.Options...)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "testdata/1/2/3/123456789098765432_720x720_fit.jpg", result)
	})
	t.Run("invalid width", func(t *testing.T) {
		colorThumb := Types["colors"]

		result, err := Filename("123456789098765432", "testdata", -2, colorThumb.Height, colorThumb.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "resample: width exceeds limit (-2)", err.Error())
		assert.Empty(t, result)
	})
	t.Run("invalid height", func(t *testing.T) {
		colorThumb := Types["colors"]

		result, err := Filename("123456789098765432", "testdata", colorThumb.Width, -3, colorThumb.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "resample: height exceeds limit (-3)", err.Error())
		assert.Empty(t, result)
	})
	t.Run("invalid hash", func(t *testing.T) {
		colorThumb := Types["colors"]

		result, err := Filename("12", "testdata", colorThumb.Width, colorThumb.Height, colorThumb.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "resample: file hash is empty or too short (12)", err.Error())
		assert.Empty(t, result)
	})
	t.Run("invalid thumb path", func(t *testing.T) {
		colorThumb := Types["colors"]

		result, err := Filename("123456789098765432", "", colorThumb.Width, colorThumb.Height, colorThumb.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "resample: folder is empty", err.Error())
		assert.Empty(t, result)
	})
}

func TestFromFile(t *testing.T) {
	t.Run("colors", func(t *testing.T) {
		colorThumb := Types["colors"]
		src := "testdata/example.gif"
		dst := "testdata/1/2/3/123456789098765432_3x3_resize.png"

		assert.FileExists(t, src)

		fileName, err := FromFile(src, "123456789098765432", "testdata", colorThumb.Width, colorThumb.Height, OrientationNormal, colorThumb.Options...)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, dst, fileName)

		assert.FileExists(t, dst)
	})

	t.Run("orientation >1 ", func(t *testing.T) {
		colorThumb := Types["colors"]
		src := "testdata/example.gif"
		dst := "testdata/1/2/3/123456789098765432_3x3_resize.png"

		assert.FileExists(t, src)

		fileName, err := FromFile(src, "123456789098765432", "testdata", colorThumb.Width, colorThumb.Height, 3, colorThumb.Options...)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, dst, fileName)

		assert.FileExists(t, dst)
	})

	t.Run("missing file", func(t *testing.T) {
		colorThumb := Types["colors"]
		src := "testdata/example.xxx"

		assert.NoFileExists(t, src)

		fileName, err := FromFile(src, "193456789098765432", "testdata", colorThumb.Width, colorThumb.Height, OrientationNormal, colorThumb.Options...)

		assert.Equal(t, "", fileName)
		assert.Error(t, err)
	})
	t.Run("empty filename", func(t *testing.T) {
		colorThumb := Types["colors"]

		fileName, err := FromFile("", "193456789098765432", "testdata", colorThumb.Width, colorThumb.Height, OrientationNormal, colorThumb.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "", fileName)
		assert.Equal(t, "resample: invalid file name “”", err.Error())
	})
}

func TestFromCache(t *testing.T) {
	t.Run("missing thumb", func(t *testing.T) {
		tile50 := Types["tile_50"]
		src := "testdata/example.jpg"

		assert.FileExists(t, src)

		fileName, err := FromCache(src, "193456789098765432", "testdata", tile50.Width, tile50.Height, tile50.Options...)

		assert.Equal(t, "", fileName)

		if err != ErrThumbNotCached {
			t.Fatal("ErrThumbNotCached expected")
		}
	})

	t.Run("missing file", func(t *testing.T) {
		tile50 := Types["tile_50"]
		src := "testdata/example.xxx"

		assert.NoFileExists(t, src)

		fileName, err := FromCache(src, "193456789098765432", "testdata", tile50.Width, tile50.Height, tile50.Options...)

		assert.Equal(t, "", fileName)
		assert.Error(t, err)
	})
	t.Run("invalid hash", func(t *testing.T) {
		tile50 := Types["tile_50"]
		src := "testdata/example.jpg"

		assert.FileExists(t, src)

		fileName, err := FromCache(src, "12", "testdata", tile50.Width, tile50.Height, tile50.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "resample: invalid file hash 12", err.Error())
		assert.Empty(t, fileName)
	})
	t.Run("empty filename", func(t *testing.T) {
		tile50 := Types["tile_50"]

		fileName, err := FromCache("", "193456789098765432", "testdata", tile50.Width, tile50.Height, tile50.Options...)

		if err == nil {
			t.Fatal("error expected")
		}
		assert.Equal(t, "resample: invalid file name “”", err.Error())
		assert.Empty(t, fileName)
	})
}

func TestCreate(t *testing.T) {
	t.Run("tile_500", func(t *testing.T) {
		tile500 := Types["tile_500"]
		src := "testdata/example.jpg"
		dst := "testdata/example.tile_500.jpg"

		assert.FileExists(t, src)
		assert.NoFileExists(t, dst)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		resized, err := Create(img, dst, tile500.Width, tile500.Height, tile500.Options...)

		if err != nil {
			t.Fatal(err)
		}

		assert.FileExists(t, dst)

		if err := os.Remove(dst); err != nil {
			t.Fatal(err)
		}

		imgNew := resized
		boundsNew := imgNew.Bounds()

		assert.Equal(t, 500, boundsNew.Max.X)
		assert.Equal(t, 500, boundsNew.Max.Y)
	})
	t.Run("width & height <= 150", func(t *testing.T) {
		tile500 := Types["tile_500"]
		src := "testdata/example.jpg"
		dst := "testdata/example.tile_500.jpg"

		assert.FileExists(t, src)
		assert.NoFileExists(t, dst)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		resized, err := Create(img, dst, 150, 150, tile500.Options...)

		if err != nil {
			t.Fatal(err)
		}

		assert.FileExists(t, dst)

		if err := os.Remove(dst); err != nil {
			t.Fatal(err)
		}

		imgNew := resized
		boundsNew := imgNew.Bounds()

		assert.Equal(t, 150, boundsNew.Max.X)
		assert.Equal(t, 150, boundsNew.Max.Y)
	})
	t.Run("invalid width", func(t *testing.T) {
		tile500 := Types["tile_500"]
		src := "testdata/example.jpg"
		dst := "testdata/example.tile_500.jpg"

		assert.FileExists(t, src)
		assert.NoFileExists(t, dst)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		resized, err := Create(img, dst, -5, tile500.Height, tile500.Options...)

		if err == nil {
			t.Fatal("error expected")
		}

		assert.Equal(t, "resample: width has an invalid value (-5)", err.Error())
		t.Log(resized)
	})
	t.Run("invalid height", func(t *testing.T) {
		tile500 := Types["tile_500"]
		src := "testdata/example.jpg"
		dst := "testdata/example.tile_500.jpg"

		assert.FileExists(t, src)
		assert.NoFileExists(t, dst)

		img, err := imaging.Open(src, imaging.AutoOrientation(true))

		if err != nil {
			t.Fatal(err)
		}

		bounds := img.Bounds()

		assert.Equal(t, 750, bounds.Max.X)
		assert.Equal(t, 500, bounds.Max.Y)

		resized, err := Create(img, dst, tile500.Width, -3, tile500.Options...)

		if err == nil {
			t.Fatal("error expected")
		}

		assert.Equal(t, "resample: height has an invalid value (-3)", err.Error())
		t.Log(resized)
	})
}
