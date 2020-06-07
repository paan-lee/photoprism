package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileName returns the a relative filename with the same base and a given extension in a directory.
func FileName(fileName, dirName, baseDir, fileExt string, stripSequence bool) string {
	fileDir := filepath.Dir(fileName)
	baseName := Base(fileName, stripSequence)

	if dirName == "" || dirName == "." {
		dirName = fileDir
	} else if fileDir != dirName {
		if filepath.IsAbs(dirName) {
			dirName = filepath.Join(dirName, Rel(fileDir, baseDir))
		} else {
			dirName = filepath.Join(fileDir, dirName)
		}
	}

	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		fmt.Println(err.Error())
		return ""
	}

	result := filepath.Join(dirName, baseName) + fileExt

	return result
}

// Rel returns the file name relative to a directory.
func Rel(fileName, dir string) string {
	if fileName == dir {
		return ""
	}

	if dir == "" {
		return fileName
	}

	if index := strings.Index(fileName, dir); index == 0 {
		if index := strings.LastIndex(dir, string(os.PathSeparator)); index == len(dir)-1 {
			pos := len(dir)
			return fileName[pos:]
		} else if index := strings.LastIndex(dir, string(os.PathSeparator)); index != len(dir) {
			pos := len(dir) + 1
			return fileName[pos:]
		}
	}

	return fileName
}
