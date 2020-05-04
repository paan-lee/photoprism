package fs

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type IgnoreLogFunc func(fileName string)

// IgnoreItem represents a file name pattern to be ignored.
type IgnoreItem struct {
	Dir     string
	Pattern string
}

// NewIgnoreItem returns a pointer to a new IgnoreItem instance.
func NewIgnoreItem(dir, pattern string, caseSensitive bool) IgnoreItem {
	if caseSensitive {
		return IgnoreItem{Dir: dir + "/", Pattern: pattern}
	} else {
		return IgnoreItem{Dir: strings.ToLower(dir) + "/", Pattern: strings.ToLower(pattern)}
	}
}

// Ignore returns true if the file name "base" in the directory "dir" should be ignored.
func (i IgnoreItem) Ignore(dir, base string) (result bool) {
	if !strings.HasPrefix(dir+"/", i.Dir) {
		// different directory prefix: don't look any further
		return false
	}

	if strings.HasPrefix(i.Pattern, "*") {
		// file name ends with pattern (e.g. ".jpg")
		result = strings.HasSuffix(base, i.Pattern[1:])
	} else if strings.HasSuffix(i.Pattern, "*") {
		// file name starts with pattern (e.g. "screenshot")
		result = strings.HasPrefix(base, i.Pattern[:len(i.Pattern)-2])
	} else if i.Pattern == base {
		// file name is the same as pattern (no wildcard)
		result = true
	}

	return result
}

// IgnoreItem represents a list of name patterns to be ignored.
type IgnoreList struct {
	Log           IgnoreLogFunc
	items         []IgnoreItem
	hiddenFiles   []string
	ignoredFiles  []string
	configFiles   map[string][]string
	configFile    string
	ignoreHidden  bool
	caseSensitive bool
}

// NewIgnoreList returns a pointer to a new IgnoreList instance.
func NewIgnoreList(configFile string, ignoreHidden bool, caseSensitive bool) *IgnoreList {
	return &IgnoreList{
		configFile:    configFile,
		ignoreHidden:  ignoreHidden,
		caseSensitive: caseSensitive,
		configFiles:   make(map[string][]string),
	}
}

// Hidden returns hidden files that were ignored.
func (l *IgnoreList) Hidden() []string {
	return l.hiddenFiles
}

// Ignored returns files that were ignored in addition to hidden files.
func (l *IgnoreList) Ignored() []string {
	return l.ignoredFiles
}

// AppendItems adds items to the list of ignored items.
func (l *IgnoreList) AppendItems(dir string, patterns []string) error {
	if dir == "" {
		return errors.New("empty directory name")
	}

	for _, pattern := range patterns {
		if pattern != "" && !strings.HasPrefix(pattern, "#") {
			l.items = append(l.items, NewIgnoreItem(dir, pattern, l.caseSensitive))
		}
	}

	return nil
}

// ConfigFile adds items in fileName to the list of ignored items.
func (l *IgnoreList) ConfigFile(fileName string) error {
	items, err := ReadLines(fileName)

	if err != nil {
		return err
	}

	l.configFiles[fileName] = items

	return l.AppendItems(filepath.Dir(fileName), items)
}

// Dir adds the ignore file in dirName to the list of ignored items.
func (l *IgnoreList) Dir(dir string) error {
	if dir == "" {
		return errors.New("empty directory name")
	}

	if l.configFile == "" {
		return errors.New("empty ignore file name")
	}

	fileName := filepath.Join(dir, l.configFile)

	if _, ok := l.configFiles[fileName]; ok {
		return nil
	}

	if !FileExists(fileName) {
		return fmt.Errorf("no %s file found", l.configFile)
	}

	return l.ConfigFile(fileName)
}

// Ignore returns true if the file name should be ignored.
func (l *IgnoreList) Ignore(fileName string) bool {
	dir := filepath.Dir(fileName)
	base := filepath.Base(fileName)

	if l.caseSensitive == false {
		dir = strings.ToLower(dir)
		base = strings.ToLower(base)
	}

	if l.configFile != "" && base == l.configFile {
		_ = l.ConfigFile(fileName)

		return true
	}

	if l.ignoreHidden && strings.HasPrefix(filepath.Base(fileName), ".") {
		l.hiddenFiles = append(l.hiddenFiles, fileName)
		return true
	}

	for _, item := range l.items {
		if item.Ignore(dir, base) {
			l.ignoredFiles = append(l.ignoredFiles, fileName)

			if l.Log != nil {
				l.Log(fileName)
			}

			return true
		}
	}

	return false
}
