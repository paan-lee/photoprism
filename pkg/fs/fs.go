/*
Package fs provides filesystem related constants and functions.

Additional information can be found in our Developer Guide:

https://github.com/photoprism/photoprism/wiki
*/
package fs

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// FileExists returns true if file exists and is not a directory.
func FileExists(fileName string) bool {
	if fileName == "" {
		return false
	}

	info, err := os.Stat(fileName)

	return err == nil && !info.IsDir()
}

// PathExists returns true if path exists and is a directory or symlink.
func PathExists(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	m := info.Mode()

	return m&os.ModeDir != 0 || m&os.ModeSymlink != 0
}

// Overwrite overwrites the file with data. Creates file if not present.
func Overwrite(fileName string, data []byte) bool {
	f, err := os.Create(fileName)
	if err != nil {
		return false
	}

	_, err = f.Write(data)
	return err == nil
}

// Abs returns the full path of a file or directory, "~" is replaced with home.
func Abs(name string) string {
	if name == "" {
		return ""
	}

	if len(name) > 2 && name[:2] == "~/" {
		if usr, err := user.Current(); err == nil {
			name = filepath.Join(usr.HomeDir, name[2:])
		}
	}

	result, err := filepath.Abs(name)

	if err != nil {
		panic(err)
	}

	return result
}

// copyToFile copies the zip file to destination
// if the zip file is a directory, a directory is created at the destination.
func copyToFile(f *zip.File, dest string) (fileName string, err error) {
	rc, err := f.Open()
	if err != nil {
		return fileName, err
	}

	defer rc.Close()

	// Store filename/path for returning and using later on
	fileName = filepath.Join(dest, f.Name)

	if f.FileInfo().IsDir() {
		// Make Folder
		return fileName, os.MkdirAll(fileName, os.ModePerm)
	}

	// Make File
	var fdir string
	if lastIndex := strings.LastIndex(fileName, string(os.PathSeparator)); lastIndex > -1 {
		fdir = fileName[:lastIndex]
	}

	err = os.MkdirAll(fdir, os.ModePerm)
	if err != nil {
		return fileName, err
	}

	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return fileName, err
	}

	defer fd.Close()
	_, err = io.Copy(fd, rc)
	if err != nil {
		return fileName, err
	}

	return fileName, nil
}

// Download downloads a file from a URL.
func Download(filepath string, url string) error {
	os.MkdirAll("/tmp/photoprism", os.ModePerm)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// IsEmpty returns true if a directory is empty.
func IsEmpty(path string) bool {
	f, err := os.Open(path)

	if err != nil {
		return false
	}

	defer f.Close()

	_, err = f.Readdirnames(1)

	if err == io.EOF {
		return true
	}

	return false
}
