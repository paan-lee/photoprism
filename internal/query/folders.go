package query

import (
	"path/filepath"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/pkg/fs"
)

// FoldersByPath returns a slice of folders in a given directory incl sub directories in recursive mode.
func FoldersByPath(rootName, rootPath, path string, recursive bool) (folders entity.Folders, err error) {
	dirs, err := fs.Dirs(filepath.Join(rootPath, path), recursive, true)

	if err != nil {
		return folders, err
	}

	folders = make(entity.Folders, len(dirs))

	for i, dir := range dirs {
		newFolder := entity.NewFolder(rootName, filepath.Join(path, dir), fs.BirthTime(filepath.Join(rootPath, dir)))

		if err := newFolder.Create(); err == nil {
			folders[i] = newFolder
		} else if folder := entity.FindFolder(rootName, filepath.Join(path, dir)); folder != nil {
			folders[i] = *folder
		} else {
			log.Errorf("folders: %s (create folder)", err)
		}
	}

	return folders, nil
}

// AlbumFolders returns folders that should be added as album.
func AlbumFolders(threshold int) (folders entity.Folders, err error) {
	db := UnscopedDb().Table("folders").
		Select("folders.path, folders.root, folders.folder_uid, folders.folder_title, folders.folder_country, folders.folder_year, folders.folder_month, COUNT(photos.id) AS photo_count").
		Joins("JOIN photos ON photos.photo_path = folders.path AND photos.deleted_at IS NULL AND photos.photo_quality >= 3").
		Group("folders.path, folders.root, folders.folder_uid, folders.folder_title, folders.folder_country, folders.folder_year, folders.folder_month").
		Having("photo_count >= ?", threshold)

	if err := db.Scan(&folders).Error; err != nil {
		return folders, err
	}

	return folders, nil
}

// UpdateFolderDates updates folder year, month and day based on indexed photo metadata.
func UpdateFolderDates() error {
	return UnscopedDb().Exec(`UPDATE folders
	INNER JOIN
	(SELECT photo_path, MAX(photo_year) AS max_photo_year, MAX(photo_month) AS max_photo_month, MAX(photo_day) AS max_photo_day
	FROM photos WHERE taken_src = 'meta' AND photos.photo_quality >= 3 AND photos.deleted_at IS NULL
	GROUP BY photo_path) AS p ON folders.path = p.photo_path
	SET folders.folder_year = p.max_photo_year, folders.folder_month = p.max_photo_month, folders.folder_day = p.max_photo_day
	WHERE p.max_photo_year IS NOT NULL AND p.max_photo_year > 0`).Error
}
