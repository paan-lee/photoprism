package photoprism

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/karrick/godirwalk"
	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/nsfw"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
)

// Index represents an indexer that indexes files in the originals directory.
type Index struct {
	conf         *config.Config
	tensorFlow   *classify.TensorFlow
	nsfwDetector *nsfw.Detector
	convert      *Convert
	db           *gorm.DB
	q            *query.Query
}

// NewIndex returns a new indexer and expects its dependencies as arguments.
func NewIndex(conf *config.Config, tensorFlow *classify.TensorFlow, nsfwDetector *nsfw.Detector, convert *Convert) *Index {
	i := &Index{
		conf:         conf,
		tensorFlow:   tensorFlow,
		nsfwDetector: nsfwDetector,
		convert:      convert,
		db:           conf.Db(),
		q:            query.New(conf.Db()),
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
	mutex.Worker.Cancel()
}

// Start indexes media files in the originals directory.
func (ind *Index) Start(opt IndexOptions) map[string]bool {
	done := make(map[string]bool)
	originalsPath := ind.originalsPath()
	optionsPath := filepath.Join(originalsPath, opt.Path)

	if !fs.PathExists(optionsPath) {
		event.Error(fmt.Sprintf("index: %s does not exist", txt.Quote(optionsPath)))
		return done
	}

	if err := mutex.Worker.Start(); err != nil {
		event.Error(fmt.Sprintf("index: %s", err.Error()))
		return done
	}

	defer mutex.Worker.Stop()

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

	ignore := fs.NewIgnoreList(IgnoreFile, true, false)

	if err := ignore.Dir(originalsPath); err != nil {
		log.Infof("index: %s", err)
	}

	ignore.Log = func(fileName string) {
		log.Infof(`index: ignored "%s"`, fs.RelativeName(fileName, originalsPath))
	}

	err := godirwalk.Walk(optionsPath, &godirwalk.Options{
		Callback: func(fileName string, info *godirwalk.Dirent) error {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("index: %s [panic]", err)
				}
			}()

			if mutex.Worker.Canceled() {
				return errors.New("indexing canceled")
			}

			isDir := info.IsDir()
			isSymlink := info.IsSymlink()

			if skip, result := fs.SkipWalk(fileName, isDir, isSymlink, done, ignore); skip {
				return result
			}

			mf, err := NewMediaFile(fileName)

			if err != nil || !mf.IsMedia() {
				return nil
			}

			related, err := mf.RelatedFiles(ind.conf.Settings().Index.Group)

			if err != nil {
				log.Warnf("index: %s", err.Error())

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

	if len(done) > 0 {
		if err := entity.UpdatePhotoCounts(); err != nil {
			log.Errorf("index: %s", err)
		}
	}

	runtime.GC()

	return done
}
