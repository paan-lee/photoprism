package fs

import (
	"net/http"
	"os"
)

const (
	MimeTypeJpeg = "image/jpeg"
)

// MimeType returns the mime type of a file, empty string if unknown.
func MimeType(filename string) string {
	handle, err := os.Open(filename)

	if err != nil {
		return ""
	}

	defer handle.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = handle.Read(buffer)

	if err != nil {
		return ""
	}

	return http.DetectContentType(buffer)
}
