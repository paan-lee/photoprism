package api

import (
	"fmt"
	"strconv"

	"github.com/photoprism/photoprism/internal/context"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/photoprism"
)

var photoIconSvg = []byte(`
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M0 0h24v24H0z" fill="none"/>
<path d="M21 19V5c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zM8.5 13.5l2.5 3.01L14.5 12l4.5 6H5l3.5-4.5z"/>
</svg>`)

// GET /api/v1/thumbnails/:type/:size/:hash
//
// Parameters:
//   type: string Format, either "fit" or "square"
//   size: int    Size in pixels
//   hash: string The file hash as returned by the search API
func GetThumbnail(router *gin.RouterGroup, ctx *context.Context) {
	router.GET("/thumbnails/:type/:size/:hash", func(c *gin.Context) {
		fileHash := c.Param("hash")
		thumbnailType := c.Param("type")
		size, err := strconv.Atoi(c.Param("size"))
		if err != nil {
			log.Errorf("invalid size: %s", c.Param("size"))
			c.Data(400, "image/svg+xml", photoIconSvg)
			return
		}

		search := photoprism.NewSearch(ctx.OriginalsPath(), ctx.Db())
		file, err := search.FindFileByHash(fileHash)

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
			return
		}

		fileName := fmt.Sprintf("%s/%s", ctx.OriginalsPath(), file.FileName)

		mediaFile, err := photoprism.NewMediaFile(fileName)
		if err != nil {
			log.Errorf("could not find image for thumbnail: %s", err.Error())
			c.Data(404, "image/svg+xml", photoIconSvg)

			// Set missing flag so that the file doesn't show up in search results anymore
			file.FileMissing = true
			ctx.Db().Save(&file)
			return
		}

		switch thumbnailType {
		case "fit":
			if thumbnail, err := mediaFile.Thumbnail(ctx.ThumbnailsPath(), size); err == nil {
				c.File(thumbnail.Filename())
			} else {
				log.Errorf("could not create thumbnail: %s", err.Error())
				c.Data(400, "image/svg+xml", photoIconSvg)
			}
		case "square":
			if thumbnail, err := mediaFile.SquareThumbnail(ctx.ThumbnailsPath(), size); err == nil {
				c.File(thumbnail.Filename())
			} else {
				log.Errorf("could not create square thumbnail: %s", err.Error())
				c.Data(400, "image/svg+xml", photoIconSvg)
			}
		default:
			log.Errorf("unknown thumbnail type: %s", thumbnailType)
			c.Data(400, "image/svg+xml", photoIconSvg)
		}
	})
}
