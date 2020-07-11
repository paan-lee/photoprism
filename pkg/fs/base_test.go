package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripExt(t *testing.T) {
	t.Run("Test.jpg", func(t *testing.T) {
		result := StripExt("/testdata/Test.jpg")
		assert.Equal(t, "/testdata/Test", result)
	})

	t.Run("Test.jpg.json", func(t *testing.T) {
		result := StripExt("/testdata/Test.jpg.json")
		assert.Equal(t, "/testdata/Test.jpg", result)
	})

	t.Run("Test copy 3.foo", func(t *testing.T) {
		result := StripExt("/testdata/Test copy 3.foo")
		assert.Equal(t, "/testdata/Test copy 3", result)
	})
}

func TestStripKnownExt(t *testing.T) {
	t.Run("Test.jpg", func(t *testing.T) {
		result := StripKnownExt("/testdata/Test.jpg")
		assert.Equal(t, "/testdata/Test", result)
	})

	t.Run("Test.jpg.json", func(t *testing.T) {
		result := StripKnownExt("/testdata/Test.jpg.json")
		assert.Equal(t, "/testdata/Test.jpg", result)
	})

	t.Run("Test copy 3.foo", func(t *testing.T) {
		result := StripKnownExt("/testdata/Test copy 3.foo")
		assert.Equal(t, "/testdata/Test copy 3.foo", result)
	})
}

func TestBase(t *testing.T) {
	t.Run("Screenshot 2019-05-21 at 10.45.52.png", func(t *testing.T) {
		regular := Base("Screenshot 2019-05-21 at 10.45.52.png", false)
		assert.Equal(t, "Screenshot 2019-05-21 at 10.45.52", regular)
		stripped := Base("Screenshot 2019-05-21 at 10.45.52.png", true)
		assert.Equal(t, "Screenshot 2019-05-21 at 10.45.52", stripped)
	})

	t.Run("Test.jpg", func(t *testing.T) {
		result := Base("/testdata/Test.jpg", true)
		assert.Equal(t, "Test", result)
	})

	t.Run("Test.jpg.json", func(t *testing.T) {
		result := Base("/testdata/Test.jpg.json", true)
		assert.Equal(t, "Test", result)
	})

	t.Run("Test copy 3.jpg", func(t *testing.T) {
		result := Base("/testdata/Test copy 3.jpg", true)
		assert.Equal(t, "Test", result)
	})

	t.Run("Test (3).jpg", func(t *testing.T) {
		result := Base("/testdata/Test (3).jpg", true)
		assert.Equal(t, "Test", result)
	})

	t.Run("Test.jpg", func(t *testing.T) {
		result := Base("/testdata/Test.jpg", false)
		assert.Equal(t, "Test", result)
	})

	t.Run("Test.3453453.jpg", func(t *testing.T) {
		regular := Base("/testdata/Test.3453453.jpg", false)
		assert.Equal(t, "Test.3453453", regular)

		stripped := Base("/testdata/Test.3453453.jpg", true)
		assert.Equal(t, "Test", stripped)
	})

	t.Run("/foo/bar.0000.ZIP", func(t *testing.T) {
		regular := Base("/foo/bar.0000.ZIP", false)
		assert.Equal(t, "bar.0000", regular)

		stripped := Base("/foo/bar.0000.ZIP", true)
		assert.Equal(t, "bar.0000", stripped)
	})

	t.Run("/foo/bar.00001.ZIP", func(t *testing.T) {
		regular := Base("/foo/bar.00001.ZIP", false)
		assert.Equal(t, "bar.00001", regular)

		stripped := Base("/foo/bar.00001.ZIP", true)
		assert.Equal(t, "bar", stripped)
	})

	t.Run("Test copy 3.jpg", func(t *testing.T) {
		result := Base("/testdata/Test copy 3.jpg", false)
		assert.Equal(t, "Test copy 3", result)
	})

	t.Run("Test (3).jpg", func(t *testing.T) {
		result := Base("/testdata/Test (3).jpg", false)
		assert.Equal(t, "Test (3)", result)
	})
	t.Run("20180506_091537_DSC02122.JPG", func(t *testing.T) {
		result := Base("20180506_091537_DSC02122.JPG", true)
		assert.Equal(t, "20180506_091537_DSC02122", result)
	})
	t.Run("20180506_091537_DSC02122 (+3.3).JPG", func(t *testing.T) {
		result := Base("20180506_091537_DSC02122 (+3.3).JPG", true)
		assert.Equal(t, "20180506_091537_DSC02122", result)
	})
	t.Run("20180506_091537_DSC02122 (-2.7).JPG", func(t *testing.T) {
		result := Base("20180506_091537_DSC02122 (-2.7).JPG", true)
		assert.Equal(t, "20180506_091537_DSC02122", result)
	})
	t.Run("20180506_091537_DSC02122(+3.3).JPG", func(t *testing.T) {
		result := Base("20180506_091537_DSC02122(+3.3).JPG", true)
		assert.Equal(t, "20180506_091537_DSC02122", result)
	})
	t.Run("20180506_091537_DSC02122(-2.7).JPG", func(t *testing.T) {
		result := Base("20180506_091537_DSC02122(-2.7).JPG", true)
		assert.Equal(t, "20180506_091537_DSC02122", result)
	})
}

func TestRelBase(t *testing.T) {
	t.Run("/foo/bar.0000.ZIP", func(t *testing.T) {
		regular := RelBase("/foo/bar.0000.ZIP", "/bar", false)
		assert.Equal(t, "/foo/bar.0000", regular)

		stripped := RelBase("/foo/bar.0000.ZIP", "/bar", true)
		assert.Equal(t, "/foo/bar.0000", stripped)
	})

	t.Run("/foo/bar.00001.ZIP", func(t *testing.T) {
		regular := RelBase("/foo/bar.00001.ZIP", "/bar", false)
		assert.Equal(t, "/foo/bar.00001", regular)

		stripped := RelBase("/foo/bar.00001.ZIP", "/bar", true)
		assert.Equal(t, "/foo/bar", stripped)
	})

	t.Run("Test copy 3.jpg", func(t *testing.T) {
		result := RelBase("/testdata/foo/Test copy 3.jpg", "/testdata", false)
		assert.Equal(t, "foo/Test copy 3", result)
	})

	t.Run("Test (3).jpg", func(t *testing.T) {
		result := RelBase("/testdata/foo/Test (3).jpg", "/testdata", false)
		assert.Equal(t, "foo/Test (3)", result)
	})

	t.Run("Test (3).jpg", func(t *testing.T) {
		result := RelBase("/testdata/foo/Test (3).jpg", "/testdata/foo/Test (3).jpg", false)
		assert.Equal(t, "Test (3)", result)
	})
}

func TestBaseAbs(t *testing.T) {
	t.Run("Test copy 3.jpg", func(t *testing.T) {
		result := AbsBase("/testdata/Test (4).jpg", true)

		assert.Equal(t, "/testdata/Test", result)
	})

	t.Run("Test (3).jpg", func(t *testing.T) {
		result := AbsBase("/testdata/Test (4).jpg", false)

		assert.Equal(t, "/testdata/Test (4)", result)
	})
}
