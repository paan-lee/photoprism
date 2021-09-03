package crop

import (
	"fmt"

	"path"

	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
)

// FromCache returns the crop file name if cached.
func FromCache(hash, thumbPath string, width, height int, area string) (fileName string, err error) {
	fileName, err = FileName(hash, thumbPath, width, height, area)

	if err != nil {
		return "", err
	}

	if fs.FileExists(fileName) {
		return fileName, nil
	}

	return "", ErrNotFound
}

// FileName returns the crop file name based on cache path, size, and area.
func FileName(hash string, thumbPath string, width, height int, area string) (fileName string, err error) {
	if len(hash) < 4 {
		return "", fmt.Errorf("crop: invalid file hash %s", txt.Quote(hash))
	}

	if len(thumbPath) < 1 {
		return "", fmt.Errorf("crop: cache path missing")
	}

	if width < 1 || height < 1 || width > 2048 || height > 2048 {
		return "", fmt.Errorf("crop: invalid size %dx%d", width, height)
	}

	fileName = path.Join(thumbPath, hash[0:1], hash[1:2], hash[2:3], fmt.Sprintf("%s_%dx%d_crop_%s%s", hash, width, height, area, fs.JpegExt))

	return fileName, nil
}
