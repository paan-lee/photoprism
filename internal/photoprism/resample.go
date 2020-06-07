package photoprism

import (
	"errors"
	"path/filepath"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/pkg/fs"
)

// Resample represents a thumbnail generator.
type Resample struct {
	conf *config.Config
}

// NewResample returns a new thumbnail generator and expects the config as argument.
func NewResample(conf *config.Config) *Resample {
	return &Resample{conf: conf}
}

// Start creates default thumbnails for all files in originalsPath.
func (rs *Resample) Start(force bool) error {
	if err := mutex.MainWorker.Start(); err != nil {
		return err
	}

	defer mutex.MainWorker.Stop()

	originalsPath := rs.conf.OriginalsPath()
	thumbnailsPath := rs.conf.ThumbPath()

	jobs := make(chan ResampleJob)

	// Start a fixed number of goroutines to read and digest files.
	var wg sync.WaitGroup
	var numWorkers = rs.conf.Workers()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			ResampleWorker(jobs)
			wg.Done()
		}()
	}

	done := make(map[string]bool)
	ignore := fs.NewIgnoreList(fs.IgnoreFile, true, false)

	if err := ignore.Dir(originalsPath); err != nil {
		log.Infof("resample: %s", err)
	}

	ignore.Log = func(fileName string) {
		log.Infof(`resample: ignored "%s"`, fs.Rel(fileName, originalsPath))
	}

	err := godirwalk.Walk(originalsPath, &godirwalk.Options{
		Callback: func(fileName string, info *godirwalk.Dirent) error {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("resample: %s [panic]", err)
				}
			}()

			if mutex.MainWorker.Canceled() {
				return errors.New("resample: canceled")
			}

			isDir := info.IsDir()
			isSymlink := info.IsSymlink()

			if skip, result := fs.SkipWalk(fileName, isDir, isSymlink, done, ignore); skip {
				return result
			}

			mf, err := NewMediaFile(fileName)

			if err != nil || !mf.IsJpeg() {
				return nil
			}

			done[fileName] = true

			relativeName := mf.RelativeName(originalsPath)

			event.Publish("index.thumbnails", event.Data{
				"fileName": relativeName,
				"baseName": filepath.Base(relativeName),
				"force":    force,
			})

			jobs <- ResampleJob{
				mediaFile: mf,
				path:      thumbnailsPath,
				force:     force,
			}

			return nil
		},
		Unsorted:            true,
		FollowSymbolicLinks: true,
	})

	close(jobs)
	wg.Wait()

	return err
}
