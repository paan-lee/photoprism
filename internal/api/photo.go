package api

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
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

		q := query.New(conf.Db())
		p, err := q.PreloadPhotoByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrPhotoNotFound)
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

		id := c.Param("uuid")
		q := query.New(conf.Db())

		m, err := q.PhotoByUUID(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrPhotoNotFound)
			return
		}

		// TODO: Proof-of-concept for form handling - might need refactoring
		// 1) Init form with model values
		f, err := form.NewPhoto(m)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		// 2) Update form with values from request
		if err := c.BindJSON(&f); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		// 3) Save model with values from form
		if err := entity.SavePhoto(m, f, conf.Db()); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		PublishPhotoEvent(EntityUpdated, id, c, q)

		event.Success("photo saved")

		p, err := q.PreloadPhotoByUUID(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrPhotoNotFound)
			return
		}

		c.JSON(http.StatusOK, p)
	})
}

// GET /api/v1/photos/:uuid/download
//
// Parameters:
//   uuid: string PhotoUUID as returned by the API
func GetPhotoDownload(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/photos/:uuid/download", func(c *gin.Context) {
		q := query.New(conf.Db())
		f, err := q.FileByPhotoUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrPhotoNotFound)
			return
		}

		fileName := path.Join(conf.OriginalsPath(), f.FileName)

		if !fs.FileExists(fileName) {
			log.Errorf("could not find original: %s", c.Param("uuid"))
			c.Data(http.StatusNotFound, "image/svg+xml", photoIconSvg)

			// Set missing flag so that the file doesn't show up in search results anymore
			f.FileMissing = true
			conf.Db().Save(&f)
			return
		}

		downloadFileName := f.ShareFileName()

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

		id := c.Param("uuid")
		q := query.New(conf.Db())
		m, err := q.PhotoByUUID(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrPhotoNotFound)
			return
		}

		m.PhotoFavorite = true
		conf.Db().Save(&m)

		event.Publish("count.favorites", event.Data{
			"count": 1,
		})

		PublishPhotoEvent(EntityUpdated, id, c, q)

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

		id := c.Param("uuid")
		q := query.New(conf.Db())
		m, err := q.PhotoByUUID(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrPhotoNotFound)
			return
		}

		m.PhotoFavorite = false
		conf.Db().Save(&m)

		event.Publish("count.favorites", event.Data{
			"count": -1,
		})

		PublishPhotoEvent(EntityUpdated, id, c, q)

		c.JSON(http.StatusOK, gin.H{"photo": m})
	})
}
