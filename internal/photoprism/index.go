package photoprism

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/photoprism/photoprism/internal/query"

	"github.com/karrick/godirwalk"
	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/nsfw"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
)

// Index represents an indexer that indexes files in the originals directory.
type Index struct {
	conf         *config.Config
	tensorFlow   *classify.TensorFlow
	nsfwDetector *nsfw.Detector
	convert      *Convert
	files        *Files
	photos       *Photos
}

// NewIndex returns a new indexer and expects its dependencies as arguments.
func NewIndex(conf *config.Config, tensorFlow *classify.TensorFlow, nsfwDetector *nsfw.Detector, convert *Convert, files *Files, photos *Photos) *Index {
	i := &Index{
		conf:         conf,
		tensorFlow:   tensorFlow,
		nsfwDetector: nsfwDetector,
		convert:      convert,
		files:        files,
		photos:       photos,
	}

	return i
}

func (ind *Index) originalsPath() string {
	return ind.conf.OriginalsPath()
}

func (ind *Index) thumbPath() string {
	return ind.conf.ThumbPath()
}

// Cancel stops the current indexing operation.
func (ind *Index) Cancel() {
	mutex.MainWorker.Cancel()
}

// Start indexes media files in the originals directory.
func (ind *Index) Start(opt IndexOptions) fs.Done {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("index: %s (panic)\nstack: %s", r, debug.Stack())
		}
	}()

	done := make(fs.Done)
	originalsPath := ind.originalsPath()
	optionsPath := filepath.Join(originalsPath, opt.Path)

	if !fs.PathExists(optionsPath) {
		event.Error(fmt.Sprintf("index: %s does not exist", txt.Quote(optionsPath)))
		return done
	}

	if err := mutex.MainWorker.Start(); err != nil {
		event.Error(fmt.Sprintf("index: %s", err.Error()))
		return done
	}

	defer mutex.MainWorker.Stop()

	if err := ind.tensorFlow.Init(); err != nil {
		log.Errorf("index: %s", err.Error())

		return done
	}

	jobs := make(chan IndexJob)

	// Start a fixed number of goroutines to index files.
	var wg sync.WaitGroup
	var numWorkers = ind.conf.Workers()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			IndexWorker(jobs) // HLc
			wg.Done()
		}()
	}

	if err := ind.files.Init(); err != nil {
		log.Errorf("index: %s", err)
	}

	defer ind.files.Done()

	filesIndexed := 0
	ignore := fs.NewIgnoreList(fs.IgnoreFile, true, false)

	if err := ignore.Dir(originalsPath); err != nil {
		log.Infof("index: %s", err)
	}

	ignore.Log = func(fileName string) {
		log.Infof(`index: ignored "%s"`, fs.RelName(fileName, originalsPath))
	}

	err := godirwalk.Walk(optionsPath, &godirwalk.Options{
		ErrorCallback: func(fileName string, err error) godirwalk.ErrorAction {
			log.Errorf("index: %s", strings.Replace(err.Error(), originalsPath, "", 1))
			return godirwalk.SkipNode
		},
		Callback: func(fileName string, info *godirwalk.Dirent) error {
			if mutex.MainWorker.Canceled() {
				return errors.New("indexing canceled")
			}

			isDir := info.IsDir()
			isSymlink := info.IsSymlink()
			relName := fs.RelName(fileName, originalsPath)

			if skip, result := fs.SkipWalk(fileName, isDir, isSymlink, done, ignore); skip {
				if (isSymlink || isDir) && result != filepath.SkipDir {
					folder := entity.NewFolder(entity.RootOriginals, relName, nil)

					if err := folder.Create(); err == nil {
						log.Infof("index: added folder /%s", folder.Path)
					}
				}

				return result
			}

			done[fileName] = fs.Found

			if !fs.IsMedia(fileName) {
				return nil
			}

			mf, err := NewMediaFile(fileName)

			if err != nil {
				log.Error(err)
				return nil
			}

			if mf.FileSize() == 0 {
				log.Infof("index: skipped empty file %s", txt.Quote(mf.BaseName()))
				return nil
			}

			if ind.files.Indexed(relName, entity.RootOriginals, mf.modTime, opt.Rescan) {
				return nil
			}

			related, err := mf.RelatedFiles(ind.conf.Settings().StackSequences())

			if err != nil {
				log.Warnf("index: %s", err.Error())

				return nil
			}

			var files MediaFiles

			for _, f := range related.Files {
				if done[f.FileName()].Processed() {
					continue
				}

				if f.FileSize() == 0 || ind.files.Indexed(f.RootRelName(), f.Root(), f.ModTime(), opt.Rescan) {
					done[f.FileName()] = fs.Found
					continue
				}

				files = append(files, f)
				filesIndexed++
				done[f.FileName()] = fs.Processed
			}

			done[fileName] = fs.Processed

			if len(files) == 0 || related.Main == nil {
				// Nothing to do.
				return nil
			}

			related.Files = files

			jobs <- IndexJob{
				FileName: mf.FileName(),
				Related:  related,
				IndexOpt: opt,
				Ind:      ind,
			}

			return nil
		},
		Unsorted:            false,
		FollowSymbolicLinks: true,
	})

	close(jobs)
	wg.Wait()

	if err != nil {
		log.Error(err.Error())
	}

	if opt.Stack {
		if err := ind.StackPhotos(); err != nil {
			log.Errorf("index: %s", err)
		}
	}

	if filesIndexed > 0 {
		if err := entity.UpdatePhotoCounts(); err != nil {
			log.Errorf("index: %s", err)
		}
	} else {
		log.Infof("index: no new or modified files")
	}

	runtime.GC()

	return done
}

// StackPhotos stacks files that belong to the same photo.
func (ind *Index) StackPhotos() error {
	photos, err := query.MatchingPhotos(ind.conf.Settings().StackMeta(), ind.conf.Settings().StackUUID())

	if err != nil {
		return err
	}

	for _, photo := range photos {
		if merged, err := photo.Stack(ind.conf.Settings().StackMeta(), ind.conf.Settings().StackUUID()); err != nil {
			log.Errorf("index: %s", err)
		} else {
			log.Infof("index: merged photo uid %s with %s", photo.PhotoUID, merged.UIDs())
			event.EntitiesUpdated("photos", []entity.Photo{photo})
			event.EntitiesDeleted("photos", merged.UIDs())
		}
	}

	return nil
}

// File indexes a single file and returns the result.
func (ind *Index) SingleFile(name string) (result IndexResult) {
	file, err := NewMediaFile(name)

	if err != nil {
		result.Err = err
		result.Status = IndexFailed

		return result
	}

	related, err := file.RelatedFiles(false)

	if err != nil {
		result.Err = err
		result.Status = IndexFailed

		return result
	}

	return IndexRelated(related, ind, IndexOptionsSingle())
}
