package photoprism

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Photo struct {
	gorm.Model
	TakenAt             time.Time
	PhotoTitle          string
	PhotoDescription    string `gorm:"type:text;"`
	PhotoArtist         string
	PhotoKeywords       string
	PhotoColors         string
	PhotoVibrantColor   string
	PhotoMutedColor     string
	PhotoCanonicalName  string
	PhotoPerceptualHash string
	PhotoFavorite       bool
	PhotoLat            float64
	PhotoLong           float64
	Location            *Location
	LocationID          uint
	Tags                []Tag `gorm:"many2many:photo_tags;"`
	Files               []File
	Albums              []Album `gorm:"many2many:album_photos;"`
	Camera              *Camera
	CameraID            uint
}
