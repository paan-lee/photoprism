package api

import (
	"net/http"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/service"
	"github.com/photoprism/photoprism/pkg/txt"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/moments/time
func GetMomentsTime(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/moments/time", func(c *gin.Context) {
		if Unauthorized(c, conf) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		q := service.Query()

		result, err := q.GetMomentsTime()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UcFirst(err.Error())})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
