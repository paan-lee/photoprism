package photoprism

import (
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/pkg/txt"
)

// Moments represents a worker that creates albums based on popular locations, dates and labels.
type Moments struct {
	conf *config.Config
}

// NewMoments returns a new purge worker.
func NewMoments(conf *config.Config) *Moments {
	instance := &Moments{
		conf: conf,
	}

	return instance
}

// Start creates albums based on popular locations, dates and categories.
func (m *Moments) Start() (err error) {
	if err := mutex.MainWorker.Start(); err != nil {
		err = fmt.Errorf("moments: %s", err.Error())
		event.Error(err.Error())
		return err
	}

	defer func() {
		mutex.MainWorker.Stop()

		if err := recover(); err != nil {
			log.Errorf("moments: %s [panic]", err)
		} else {
			runtime.GC()
		}
	}()

	counts := query.Counts{}
	counts.Refresh()

	indexSize := counts.Photos + counts.Videos

	threshold := int(math.Log2(float64(indexSize))) + 1

	log.Infof("moments: index contains %d photos and videos, threshold %d", indexSize, threshold)

	// Important folders.
	if results, err := query.AlbumFolders(threshold); err != nil {
		log.Errorf("moments: %s", err.Error())
	} else {
		for _, mom := range results {
			f := form.PhotoSearch{
				Path: mom.Path,
			}

			if a := entity.FindAlbum(mom.Slug(), entity.TypeFolder); a != nil {
				log.Infof("moments: %s already exists (%s)", txt.Quote(mom.Title()), f.Serialize())
			} else if a := entity.NewFolderAlbum(mom.Title(), mom.Slug(), f.Serialize()); a != nil {
				if err := a.Create(); err != nil {
					log.Errorf("moments: %s", err)
				} else {
					log.Infof("moments: added %s (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
				}
			}
		}
	}

	// Important years and months.
	if results, err := query.MomentsTime(threshold); err != nil {
		log.Errorf("moments: %s", err.Error())
	} else {
		for _, mom := range results {
			if a := entity.FindAlbum(mom.Slug(), entity.TypeMonth); a != nil {
				log.Infof("moments: %s already exists (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
			} else if a := entity.NewMonthAlbum(mom.Title(), mom.Slug(), mom.Year, mom.Month); a != nil {
				if err := a.Create(); err != nil {
					log.Errorf("moments: %s", err)
				} else {
					log.Infof("moments: added %s (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
				}
			}
		}
	}

	// Countries by year.
	if results, err := query.MomentsCountries(threshold); err != nil {
		log.Errorf("moments: %s", err.Error())
	} else {
		for _, mom := range results {
			f := form.PhotoSearch{
				Country: mom.Country,
				Year:    mom.Year,
			}

			if a := entity.FindAlbum(mom.Slug(), entity.TypeMoment); a != nil {
				log.Infof("moments: %s already exists (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
			} else if a := entity.NewMomentsAlbum(mom.Title(), mom.Slug(), f.Serialize()); a != nil {
				a.AlbumYear = mom.Year
				a.AlbumCountry = mom.Country

				if err := a.Create(); err != nil {
					log.Errorf("moments: %s", err)
				} else {
					log.Infof("moments: added %s (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
				}
			}
		}
	}

	// States and countries.
	if results, err := query.MomentsStates(threshold); err != nil {
		log.Errorf("moments: %s", err.Error())
	} else {
		for _, mom := range results {
			f := form.PhotoSearch{
				Country: mom.Country,
				State:   mom.State,
			}

			if a := entity.FindAlbum(mom.Slug(), entity.TypeMoment); a != nil {
				log.Infof("moments: %s already exists (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
			} else if a := entity.NewMomentsAlbum(mom.Title(), mom.Slug(), f.Serialize()); a != nil {
				a.AlbumCountry = mom.Country

				if err := a.Create(); err != nil {
					log.Errorf("moments: %s", err)
				} else {
					log.Infof("moments: added %s (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
				}
			}
		}
	}

	// Popular labels.
	if results, err := query.MomentsLabels(threshold); err != nil {
		log.Errorf("moments: %s", err.Error())
	} else {
		for _, mom := range results {
			f := form.PhotoSearch{
				Label: mom.Label,
			}

			if a := entity.FindAlbum(mom.Slug(), entity.TypeMoment); a != nil {
				log.Infof("moments: %s already exists (%s)", txt.Quote(mom.Title()), f.Serialize())

				if err := form.ParseQueryString(&f); err != nil {
					log.Errorf("moments: %s", err.Error())
				} else {
					w := txt.Words(f.Label)
					w = append(w, mom.Label)
					f.Label = strings.Join(txt.UniqueWords(w), ",")
				}

				if err := a.Update("AlbumFilter", f.Serialize()); err != nil {
					log.Errorf("moments: %s", err.Error())
				} else {
					log.Infof("moments: updated %s (%s)", txt.Quote(a.AlbumTitle), f.Serialize())
				}
			} else if a := entity.NewMomentsAlbum(mom.Title(), mom.Slug(), f.Serialize()); a != nil {
				if err := a.Create(); err != nil {
					log.Errorf("moments: %s", err.Error())
				} else {
					log.Infof("moments: added %s (%s)", txt.Quote(a.AlbumTitle), a.AlbumFilter)
				}
			} else {
				log.Errorf("moments: failed to create new moment %s (%s)", mom.Title(), f.Serialize())
			}
		}
	}

	return nil
}

// Cancel stops the current operation.
func (m *Moments) Cancel() {
	mutex.MainWorker.Cancel()
}
