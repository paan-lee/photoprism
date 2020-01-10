package api

import (
	"fmt"
	"net/http"
	"path"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/file"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/internal/txt"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/photos/:uuid
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func GetPhoto(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/photos/:uuid", func(c *gin.Context) {
		if Unauthorized(c, conf) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		q := query.New(conf.OriginalsPath(), conf.Db())
		p, err := q.PreloadPhotoByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	})
}

// PUT /api/v1/photos/:uuid
func UpdatePhoto(router *gin.RouterGroup, conf *config.Config) {
	router.PUT("/photos/:uuid", func(c *gin.Context) {
		if Unauthorized(c, conf) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		q := query.New(conf.OriginalsPath(), conf.Db())

		m, err := q.FindPhotoByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		if err := c.BindJSON(&m); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		conf.Db().Save(&m)

		event.Success("photo saved")

		c.JSON(http.StatusOK, m)
	})
}

// GET /api/v1/photos/:uuid/download
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func GetPhotoDownload(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/photos/:uuid/download", func(c *gin.Context) {
		q := query.New(conf.OriginalsPath(), conf.Db())
		f, err := q.FindFileByPhotoUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
			return
		}

		fileName := path.Join(conf.OriginalsPath(), f.FileName)

		if !file.Exists(fileName) {
			log.Errorf("could not find original: %s", c.Param("uuid"))
			c.Data(404, "image/svg+xml", photoIconSvg)

			// Set missing flag so that the file doesn't show up in search results anymore
			f.FileMissing = true
			conf.Db().Save(&f)
			return
		}

		downloadFileName := f.DownloadFileName()

		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", downloadFileName))

		c.File(fileName)
	})
}

// POST /api/v1/photos/:uuid/like
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func LikePhoto(router *gin.RouterGroup, conf *config.Config) {
	router.POST("/photos/:uuid/like", func(c *gin.Context) {
		if Unauthorized(c, conf) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		q := query.New(conf.OriginalsPath(), conf.Db())
		m, err := q.FindPhotoByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		m.PhotoFavorite = true
		conf.Db().Save(&m)

		event.Publish("count.favorites", event.Data{
			"count": 1,
		})

		c.JSON(http.StatusOK, gin.H{"photo": m})
	})
}

// DELETE /api/v1/photos/:uuid/like
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func DislikePhoto(router *gin.RouterGroup, conf *config.Config) {
	router.DELETE("/photos/:uuid/like", func(c *gin.Context) {
		if Unauthorized(c, conf) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		q := query.New(conf.OriginalsPath(), conf.Db())
		m, err := q.FindPhotoByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		m.PhotoFavorite = false
		conf.Db().Save(&m)

		event.Publish("count.favorites", event.Data{
			"count": -1,
		})

		c.JSON(http.StatusOK, gin.H{"photo": m})
	})
}
