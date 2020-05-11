package entity

import (
	"strings"
	"time"

	"github.com/photoprism/photoprism/pkg/txt"
)

var QualityBlacklist = map[string]bool{
	"screenshot":  true,
	"screenshots": true,
	"info":        true,
}

var (
	year2008 = time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)
	year2012 = time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
)

// QualityScore returns a score based on photo properties like size and metadata.
func (m *Photo) QualityScore() (score int) {
	if m.PhotoFavorite {
		score += 3
	}

	if m.TakenSrc != SrcAuto {
		score++
	}

	if m.HasLatLng() {
		score++
	}

	if m.TakenAt.Before(year2008) {
		score++
	} else if m.TakenAt.Before(year2012) && m.PhotoResolution >= 1 {
		score++
	} else if m.PhotoResolution >= 2 {
		score++
	}

	blacklisted := false

	if m.Description.PhotoKeywords != "" {
		keywords := txt.Words(m.Description.PhotoKeywords)

		for _, w := range keywords {
			w = strings.ToLower(w)

			if _, ok := QualityBlacklist[w]; ok {
				blacklisted = true
				break
			}
		}
	}

	if !blacklisted {
		score++
	}

	if score < 3 && (m.PhotoVideo || m.EditedAt != nil) {
		score = 3
	}

	return score
}
