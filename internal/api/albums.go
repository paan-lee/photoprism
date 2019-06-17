package api

import (
	"github.com/photoprism/photoprism/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/forms"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/util"

	log "github.com/sirupsen/logrus"
)

// GET /api/v1/albums
func GetAlbums(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/albums", func(c *gin.Context) {
		var form forms.AlbumSearchForm

		search := photoprism.NewSearch(conf.OriginalsPath(), conf.Db())
		err := c.MustBindWith(&form, binding.Form)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.UcFirst(err.Error())})
			return
		}

		result, err := search.Albums(form)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": util.UcFirst(err.Error())})
			return
		}

		c.Header("x-result-count", strconv.Itoa(form.Count))
		c.Header("x-result-offset", strconv.Itoa(form.Offset))

		c.JSON(http.StatusOK, result)
	})
}

type CreateAlbumParams struct {
	AlbumName string `json:"AlbumName"`
}

// POST /api/v1/albums
func CreateAlbum(router *gin.RouterGroup, conf *config.Config) {
	router.POST("/albums", func(c *gin.Context) {
		var params CreateAlbumParams

		if err := c.BindJSON(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.UcFirst(err.Error())})
		}

		if len(params.AlbumName) == 0 {
			log.Error("album name empty")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.UcFirst("album name empty")})
		}

		album := &models.Album{AlbumName: params.AlbumName}

		if res := conf.Db().Create(album); res.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.UcFirst(res.Error.Error())})
		}

		c.JSON(http.StatusOK, http.Response{})
	})
}

// POST /api/v1/albums/:uuid/like
//
// Parameters:
//   uuid: string Album UUID
func LikeAlbum(router *gin.RouterGroup, conf *config.Config) {
	router.POST("/albums/:uuid/like", func(c *gin.Context) {
		search := photoprism.NewSearch(conf.OriginalsPath(), conf.Db())

		album, err := search.FindAlbumByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": util.UcFirst(err.Error())})
			return
		}

		album.AlbumFavorite = true
		conf.Db().Save(&album)

		c.JSON(http.StatusOK, http.Response{})
	})
}

// DELETE /api/v1/albums/:uuid/like
//
// Parameters:
//   uuid: string Album UUID
func DislikeAlbum(router *gin.RouterGroup, conf *config.Config) {
	router.DELETE("/albums/:uuid/like", func(c *gin.Context) {
		search := photoprism.NewSearch(conf.OriginalsPath(), conf.Db())

		album, err := search.FindAlbumByUUID(c.Param("uuid"))

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": util.UcFirst(err.Error())})
			return
		}

		album.AlbumFavorite = false
		conf.Db().Save(&album)

		c.JSON(http.StatusOK, http.Response{})
	})
}
