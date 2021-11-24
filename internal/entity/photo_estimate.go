package entity

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/geo"
	"github.com/photoprism/photoprism/pkg/txt"
)

// EstimateCountry updates the photo with an estimated country if possible.
func (m *Photo) EstimateCountry() {
	if SrcPriority[m.PlaceSrc] > SrcPriority[SrcEstimate] || m.HasLocation() || m.HasPlace() {
		// Ignore.
		return
	}

	// Reset country.
	unknown := UnknownCountry.ID
	countryCode := unknown

	// Try to guess country from photo title.
	if code := txt.CountryCode(m.PhotoTitle); code != unknown && m.TitleSrc != SrcAuto {
		countryCode = code
	}

	// Try to guess country from filename and path.
	if countryCode == unknown {
		if code := txt.CountryCode(m.PhotoName); code != unknown && !fs.IsGenerated(m.PhotoName) {
			countryCode = code
		} else if code := txt.CountryCode(m.PhotoPath); code != unknown {
			countryCode = code
		}
	}

	// Try to guess country from original filename.
	if countryCode == unknown && m.OriginalName != "" && !fs.IsGenerated(m.OriginalName) {
		if code := txt.CountryCode(m.OriginalName); code != UnknownCountry.ID {
			countryCode = code
		}
	}

	// Set new country?
	if countryCode != unknown {
		m.PhotoCountry = countryCode
		m.PlaceSrc = SrcEstimate
		m.EstimatedAt = TimePointer()
		log.Debugf("photo: estimated country for %s is %s", m, txt.Quote(m.CountryName()))
	}
}

// EstimateLocation updates the photo with an estimated place and country if possible.
func (m *Photo) EstimateLocation(force bool) {
	if SrcPriority[m.PlaceSrc] > SrcPriority[SrcEstimate] || m.HasLocation() && m.PlaceSrc == SrcAuto {
		// Ignore if location was set otherwise.
		return
	} else if force || m.EstimatedAt == nil {
		// Proceed.
	} else if hours := TimeStamp().Sub(*m.EstimatedAt) / time.Hour; hours < MetadataEstimateInterval {
		// Ignore if estimated less than 7 days ago.
		return
	}

	// Estimate country if taken date is unreliable.
	if SrcPriority[m.TakenSrc] <= SrcPriority[SrcName] {
		m.RemoveLocation(false)
		m.EstimateCountry()
		return
	}

	var err error

	rangeMin := m.TakenAt.Add(-1 * time.Hour * 37)
	rangeMax := m.TakenAt.Add(time.Hour * 37)

	// Find photo with location info taken at a similar time...
	var mostRecent Photos

	switch DbDialect() {
	case MySQL:
		err = UnscopedDb().
			Where("photo_lat <> 0 AND photo_lng <> 0").
			Where("place_src <> '' AND place_src <> ? AND place_id IS NOT NULL AND place_id <> '' AND place_id <> 'zz'", SrcEstimate).
			Where("taken_src <> '' AND taken_at BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)", rangeMin, rangeMax).
			Order(gorm.Expr("ABS(TIMESTAMPDIFF(SECOND, taken_at, ?))", m.TakenAt)).Limit(2).
			Preload("Place").Find(&mostRecent).Error
	case SQLite:
		err = UnscopedDb().
			Where("photo_lat <> 0 AND photo_lng <> 0").
			Where("place_src <> '' AND place_src <> ? AND place_id IS NOT NULL AND place_id <> '' AND place_id <> 'zz'", SrcEstimate).
			Where("taken_src <> '' AND taken_at BETWEEN ? AND ?", rangeMin, rangeMax).
			Order(gorm.Expr("ABS(JulianDay(taken_at) - JulianDay(?))", m.TakenAt)).Limit(2).
			Preload("Place").Find(&mostRecent).Error
	default:
		log.Warnf("photo: unsupported sql dialect %s", txt.Quote(DbDialect()))
		return
	}

	if err != nil {
		log.Warnf("photo: %s while estimating position", err)
	}

	// Found?
	if len(mostRecent) == 0 {
		log.Debugf("photo: unknown position at %s", m.TakenAt)
		m.RemoveLocation(false)
		m.EstimateCountry()
	} else if recentPhoto := mostRecent[0]; recentPhoto.HasLocation() && recentPhoto.HasPlace() {
		// Too much time difference?
		if hours := recentPhoto.TakenAt.Sub(m.TakenAt) / time.Hour; hours < -36 || hours > 36 {
			log.Debugf("photo: skipping %s, %d hours time difference to recent position", m, hours)
			m.RemoveLocation(false)
			m.EstimateCountry()
		} else if len(mostRecent) == 1 {
			m.RemoveLocation(false)

			m.Place = recentPhoto.Place
			m.PlaceID = recentPhoto.PlaceID
			m.PhotoCountry = recentPhoto.PhotoCountry
			m.PlaceSrc = SrcEstimate

			m.UpdateTimeZone(recentPhoto.TimeZone)

			log.Debugf("photo: approximate place of %s is %s (id %s)", m, txt.Quote(m.Place.Label()), recentPhoto.PlaceID)
		} else if recentPhoto.HasPlace() {
			p1 := mostRecent[0]
			p2 := mostRecent[1]

			movement := geo.NewMovement(p1.Position(), p2.Position())

			if movement.Km() < 100 {
				estimate := movement.EstimatePosition(m.TakenAt)

				if m.CellID != UnknownID && estimate.InRange(float64(m.PhotoLat), float64(m.PhotoLng), geo.Meter*50) {
					log.Debugf("photo: keeping position estimate %f, %f for %s", m.PhotoLat, m.PhotoLng, m.String())
				} else {
					estimate.Randomize(geo.Meter * 5)

					m.PhotoLat = float32(estimate.Lat)
					m.PhotoLng = float32(estimate.Lng)
					m.PlaceSrc = SrcEstimate
					m.CellAccuracy = estimate.Accuracy
					m.SetAltitude(estimate.AltitudeInt(), SrcEstimate)

					log.Debugf("photo: %s %s", m.String(), estimate.String())

					m.UpdateLocation()

					if m.Place == nil {
						log.Warnf("photo: failed updating position of %s", m)
					} else {
						log.Debugf("photo: approximate place of %s is %s (id %s)", m, txt.Quote(m.Place.Label()), m.PlaceID)
					}
				}
			} else {
				m.RemoveLocation(false)

				m.Place = recentPhoto.Place
				m.PlaceID = recentPhoto.PlaceID
				m.PlaceSrc = SrcEstimate
				m.PhotoCountry = recentPhoto.PhotoCountry
				m.SetAltitude(movement.EstimateAltitudeInt(m.TakenAt), SrcEstimate)

				m.UpdateTimeZone(recentPhoto.TimeZone)
			}
		} else if recentPhoto.HasCountry() {
			m.RemoveLocation(false)
			m.PhotoCountry = recentPhoto.PhotoCountry
			m.PlaceSrc = SrcEstimate
			m.UpdateTimeZone(recentPhoto.TimeZone)

			log.Debugf("photo: probable country for %s is %s", m, txt.Quote(m.CountryName()))
		} else {
			m.RemoveLocation(false)
			m.EstimateCountry()
		}
	} else {
		log.Warnf("photo: %s has no location, uid %s", recentPhoto.PhotoName, recentPhoto.PhotoUID)
		m.RemoveLocation(false)
		m.EstimateCountry()
	}

	m.EstimatedAt = TimePointer()
}
