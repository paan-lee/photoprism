package photoprism

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
)

// Import represents an importer that can copy/move MediaFiles to the originals directory.
type Import struct {
	conf    *config.Config
	index   *Index
	convert *Convert
}

// NewImport returns a new importer and expects its dependencies as arguments.
func NewImport(conf *config.Config, index *Index, convert *Convert) *Import {
	instance := &Import{
		conf:    conf,
		index:   index,
		convert: convert,
	}

	return instance
}

// originalsPath returns the original media files path as string.
func (imp *Import) originalsPath() string {
	return imp.conf.OriginalsPath()
}

// thumbPath returns the thumbnails cache path as string.
func (imp *Import) thumbPath() string {
	return imp.conf.ThumbPath()
}

// Start imports media files from a directory and converts/indexes them as needed.
func (imp *Import) Start(opt ImportOptions) map[string]bool {
	var directories []string
	done := make(map[string]bool)
	ind := imp.index
	importPath := opt.Path

	if !fs.PathExists(importPath) {
		event.Error(fmt.Sprintf("import: %s does not exist", importPath))
		return done
	}

	if err := mutex.MainWorker.Start(); err != nil {
		event.Error(fmt.Sprintf("import: %s", err.Error()))
		return done
	}

	defer mutex.MainWorker.Stop()

	if err := ind.tensorFlow.Init(); err != nil {
		log.Errorf("import: %s", err.Error())
		return done
	}

	jobs := make(chan ImportJob)

	// Start a fixed number of goroutines to import files.
	var wg sync.WaitGroup
	var numWorkers = ind.conf.Workers()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			ImportWorker(jobs)
			wg.Done()
		}()
	}

	indexOpt := IndexOptionsAll()
	ignore := fs.NewIgnoreList(fs.IgnoreFile, true, false)

	if err := ignore.Dir(importPath); err != nil {
		log.Infof("import: %s", err)
	}

	ignore.Log = func(fileName string) {
		log.Infof(`import: ignored "%s"`, fs.RelativeName(fileName, importPath))
	}

	err := godirwalk.Walk(importPath, &godirwalk.Options{
		Callback: func(fileName string, info *godirwalk.Dirent) error {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("import: %s [panic]", err)
				}
			}()

			if mutex.MainWorker.Canceled() {
				return errors.New("import canceled")
			}

			isDir := info.IsDir()
			isSymlink := info.IsSymlink()

			if isDir {
				if fileName != importPath {
					directories = append(directories, fileName)
				}
			}

			if skip, result := fs.SkipWalk(fileName, isDir, isSymlink, done, ignore); skip {
				if isDir && result != filepath.SkipDir {
					folder := entity.NewFolder(entity.RootImport, fs.RelativeName(fileName, imp.conf.ImportPath()), nil)

					if err := folder.Create(); err == nil {
						log.Infof("import: added folder /%s", folder.Path)
					}
				}

				return result
			}

			mf, err := NewMediaFile(fileName)

			if err != nil || !mf.IsMedia() {
				return nil
			}

			related, err := mf.RelatedFiles(imp.conf.Settings().Index.Group)

			if err != nil {
				event.Error(fmt.Sprintf("import: %s", err.Error()))

				return nil
			}

			var files MediaFiles

			for _, f := range related.Files {
				if done[f.FileName()] {
					continue
				}

				files = append(files, f)
				done[f.FileName()] = true
			}

			done[fileName] = true

			related.Files = files

			jobs <- ImportJob{
				FileName:  fileName,
				Related:   related,
				IndexOpt:  indexOpt,
				ImportOpt: opt,
				Imp:       imp,
			}

			return nil
		},
		Unsorted:            false,
		FollowSymbolicLinks: true,
	})

	close(jobs)
	wg.Wait()

	sort.Slice(directories, func(i, j int) bool {
		return len(directories[i]) > len(directories[j])
	})

	if opt.RemoveEmptyDirectories {
		// Remove empty directories from import path
		for _, directory := range directories {
			if fs.IsEmpty(directory) {
				if err := os.Remove(directory); err != nil {
					log.Errorf("import: could not delete empty folder %s (%s)", txt.Quote(fs.RelativeName(directory, importPath)), err)
				} else {
					log.Infof("import: deleted empty folder %s", txt.Quote(fs.RelativeName(directory, importPath)))
				}
			}
		}
	}

	if opt.RemoveDotFiles {
		// Remove hidden .files if option is enabled
		for _, file := range ignore.Hidden() {
			if !fs.FileExists(file) {
				continue
			}

			if err := os.Remove(file); err != nil {
				log.Errorf("import: could not remove %s (%s)", txt.Quote(fs.RelativeName(file, importPath)), err.Error())
			}
		}
	}

	if err != nil {
		log.Error(err.Error())
	}

	if len(done) > 0 {
		if err := entity.UpdatePhotoCounts(); err != nil {
			log.Errorf("import: %s", err)
		}
	}

	runtime.GC()

	return done
}

// Cancel stops the current import operation.
func (imp *Import) Cancel() {
	mutex.MainWorker.Cancel()
}

// DestinationFilename returns the destination filename of a MediaFile to be imported.
func (imp *Import) DestinationFilename(mainFile *MediaFile, mediaFile *MediaFile) (string, error) {
	fileName := mainFile.CanonicalName()
	fileExtension := mediaFile.Extension()
	dateCreated := mainFile.DateCreated()

	if !mediaFile.IsSidecar() {
		if f, err := entity.FirstFileByHash(mediaFile.Hash()); err == nil {
			existingFilename := filepath.Join(imp.conf.OriginalsPath(), f.FileName)
			if fs.FileExists(existingFilename) {
				return existingFilename, fmt.Errorf("%s is identical to %s (sha1 %s)", txt.Quote(filepath.Base(mediaFile.FileName())), txt.Quote(f.FileName), mediaFile.Hash())
			} else {
				return existingFilename, nil
			}
		}
	}

	//	Mon Jan 2 15:04:05 -0700 MST 2006
	pathName := filepath.Join(imp.originalsPath(), dateCreated.Format("2006/01"))

	iteration := 0

	result := filepath.Join(pathName, fileName+fileExtension)

	for fs.FileExists(result) {
		if mediaFile.Hash() == fs.Hash(result) {
			return result, fmt.Errorf("%s already exists", txt.Quote(fs.RelativeName(result, imp.originalsPath())))
		}

		iteration++

		result = filepath.Join(pathName, fileName+"."+fmt.Sprintf("%05d", iteration)+fileExtension)
	}

	return result, nil
}
