package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/api"
	"github.com/photoprism/photoprism/internal/config"
)

func registerRoutes(router *gin.Engine, conf *config.Config) {
	// Static favicon file.
	router.StaticFile("/favicon.ico", conf.FaviconsPath()+"/favicon.ico")

	// Other static assets like JS and CSS files.
	router.Static("/static", conf.StaticPath())

	// Rainbow page.
	router.GET("/rainbow", func(c *gin.Context) {
		clientConfig := conf.PublicConfig()
		c.HTML(http.StatusOK, "rainbow.tmpl", gin.H{"config": clientConfig})
	})

	// JSON-REST API Version 1
	v1 := router.Group("/api/v1")
	{
		api.GetStatus(v1)
		api.GetConfig(v1)

		api.CreateSession(v1)
		api.DeleteSession(v1)

		api.GetPreview(v1)
		api.GetThumbnail(v1)
		api.GetDownload(v1)
		api.GetVideo(v1)
		api.CreateZip(v1)
		api.DownloadZip(v1)

		api.GetGeo(v1)
		api.GetPhoto(v1)
		api.GetPhotoYaml(v1)
		api.UpdatePhoto(v1)
		api.GetPhotos(v1)
		api.GetPhotoDownload(v1)
		api.GetPhotoLinks(v1)
		api.CreatePhotoLink(v1)
		api.UpdatePhotoLink(v1)
		api.DeletePhotoLink(v1)
		api.ApprovePhoto(v1)
		api.LikePhoto(v1)
		api.DislikePhoto(v1)
		api.AddPhotoLabel(v1)
		api.RemovePhotoLabel(v1)
		api.UpdatePhotoLabel(v1)
		api.GetMomentsTime(v1)
		api.GetFile(v1)
		api.SetPhotoPrimary(v1)

		api.GetLabels(v1)
		api.UpdateLabel(v1)
		api.GetLabelLinks(v1)
		api.CreateLabelLink(v1)
		api.UpdateLabelLink(v1)
		api.DeleteLabelLink(v1)
		api.LikeLabel(v1)
		api.DislikeLabel(v1)
		api.LabelThumbnail(v1)

		api.GetFoldersOriginals(v1)
		api.GetFoldersImport(v1)

		api.Upload(v1)
		api.StartImport(v1)
		api.CancelImport(v1)
		api.StartIndexing(v1)
		api.CancelIndexing(v1)

		api.BatchPhotosArchive(v1)
		api.BatchPhotosRestore(v1)
		api.BatchPhotosPrivate(v1)
		api.BatchAlbumsDelete(v1)
		api.BatchLabelsDelete(v1)

		api.GetAlbum(v1)
		api.CreateAlbum(v1)
		api.UpdateAlbum(v1)
		api.DeleteAlbum(v1)
		api.DownloadAlbum(v1)
		api.GetAlbums(v1)
		api.GetAlbumLinks(v1)
		api.CreateAlbumLink(v1)
		api.UpdateAlbumLink(v1)
		api.DeleteAlbumLink(v1)
		api.LikeAlbum(v1)
		api.DislikeAlbum(v1)
		api.AlbumThumbnail(v1)
		api.CloneAlbums(v1)
		api.AddPhotosToAlbum(v1)
		api.RemovePhotosFromAlbum(v1)

		api.GetAccounts(v1)
		api.GetAccount(v1)
		api.GetAccountDirs(v1)
		api.ShareWithAccount(v1)
		api.CreateAccount(v1)
		api.DeleteAccount(v1)
		api.UpdateAccount(v1)

		api.GetSettings(v1)
		api.SaveSettings(v1)

		api.GetSvg(v1)

		api.Websocket(v1)
	}

	// Link sharing.
	share := router.Group("/s")
	{
		api.InitShare(share)
	}

	// WebDAV server for file management, sync and sharing.
	if conf.WebDAVPassword() != "" {
		log.Info("webdav: enabled, username: photoprism")

		WebDAV(conf.OriginalsPath(), router.Group("/originals", gin.BasicAuth(gin.Accounts{
			"photoprism": conf.WebDAVPassword(),
		})), conf)

		log.Info("webdav: /originals/ available")

		if conf.ReadOnly() {
			log.Info("webdav: /import/ not available in read-only mode")
		} else {
			WebDAV(conf.ImportPath(), router.Group("/import", gin.BasicAuth(gin.Accounts{
				"photoprism": conf.WebDAVPassword(),
			})), conf)

			log.Info("webdav: /import/ available")
		}
	} else {
		log.Info("webdav: disabled (no password set)")
	}

	// Default HTML page for client-side rendering and routing via VueJS.
	router.NoRoute(func(c *gin.Context) {
		clientConfig := conf.PublicConfig()
		c.HTML(http.StatusOK, conf.DefaultTemplate(), gin.H{"config": clientConfig})
	})
}
