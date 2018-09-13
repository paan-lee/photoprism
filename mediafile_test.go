package photoprism

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediaFile_FindRelatedImages(t *testing.T) {
	conf := NewTestConfig()

	conf.InitializeTestData(t)

	mediaFile, err := NewMediaFile(conf.ImportPath + "/raw/20140717_154212_1EC48F8489.cr2")

	assert.Nil(t, err)

	expectedBaseFilename := conf.ImportPath + "/raw/20140717_154212_1EC48F8489"

	related, _, err := mediaFile.GetRelatedFiles()

	assert.Empty(t, err)

	assert.Len(t, related, 3)

	for _, result := range related {
		filename := result.GetFilename()

		extension := result.GetExtension()

		baseFilename := filename[0 : len(filename)-len(extension)]

		assert.Equal(t, expectedBaseFilename, baseFilename)
	}
}

func TestMediaFile_GetEditedFilename(t *testing.T) {
	conf := NewTestConfig()

	conf.InitializeTestData(t)

	mediaFile1, err := NewMediaFile(conf.ImportPath + "/iphone/IMG_6788.jpg")
	assert.Nil(t, err)
	assert.Equal(t, conf.ImportPath+"/iphone/IMG_E6788.jpg", mediaFile1.GetEditedFilename())

	/* TODO: Add example files to import.zip
	mediaFile2, err := NewMediaFile("/foo/bar/IMG_E1234.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "", mediaFile2.GetEditedFilename())
	*/

	mediaFile3, err := NewMediaFile(conf.ImportPath + "/raw/20140717_154212_1EC48F8489.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "", mediaFile3.GetEditedFilename())
}

func TestMediaFile_GetPerceptiveHash(t *testing.T) {
	conf := NewTestConfig()

	conf.InitializeTestData(t)

	mediaFile1, err := NewMediaFile(conf.ImportPath + "/20130203_193332_0AE340D280.jpg")
	assert.Nil(t, err)
	hash1, _ := mediaFile1.GetPerceptualHash()

	assert.Equal(t, "66debc383325d3bd", hash1)

	mediaFile2, err := NewMediaFile(conf.ImportPath + "/20130203_193332_0AE340D280_V2.jpg")
	assert.Nil(t, err)
	hash2, _ := mediaFile2.GetPerceptualHash()

	assert.Equal(t, "e6debc393325c3b9", hash2)

	distance, _ := mediaFile1.GetPerceptualDistance(hash2)

	assert.Equal(t, 4, distance)

	mediaFile3, err := NewMediaFile(conf.ImportPath + "/iphone/IMG_6788.JPG")
	assert.Nil(t, err)
	hash3, _ := mediaFile3.GetPerceptualHash()

	assert.Equal(t, "f1e2858b171d3e78", hash3)

	distance, _ = mediaFile1.GetPerceptualDistance(hash3)

	assert.Equal(t, 33, distance)
}

func TestMediaFile_GetMimeType(t *testing.T) {
	conf := NewTestConfig()

	conf.InitializeTestData(t)

	image1, err := NewMediaFile(conf.ImportPath + "/iphone/IMG_6788.JPG")
	assert.Nil(t, err)
	assert.Equal(t, "image/jpeg", image1.GetMimeType())

	image2, err := NewMediaFile(conf.ImportPath + "/raw/20140717_154212_1EC48F8489.cr2")
	assert.Nil(t, err)
	assert.Equal(t, "application/octet-stream", image2.GetMimeType())
}

func TestMediaFile_Exists(t *testing.T) {
	conf := NewTestConfig()

	mediaFile, err := NewMediaFile(conf.ImportPath + "/iphone/IMG_6788.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, mediaFile)
	assert.True(t, mediaFile.Exists())

	mediaFile, err = NewMediaFile(conf.ImportPath + "/iphone/IMG_6788_XYZ.JPG")
	assert.NotNil(t, err)
	assert.Nil(t, mediaFile)
}
