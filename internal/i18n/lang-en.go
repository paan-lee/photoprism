package i18n

const (
	ErrUnexpected Message = iota + 1
	ErrBadRequest
	ErrSaveFailed
	ErrDeleteFailed
	ErrAlreadyExists
	ErrNotFound
	ErrFileNotFound
	ErrSelectionNotFound
	ErrEntityNotFound
	ErrAccountNotFound
	ErrUserNotFound
	ErrLabelNotFound
	ErrAlbumNotFound
	ErrPublic
	ErrReadOnly
	ErrUnauthorized
	ErrOffensiveUpload
	ErrNoItemsSelected
	ErrCreateFile
	ErrCreateFolder
	ErrConnectionFailed
	ErrInvalidPassword
	ErrFeatureDisabled
	ErrNoLabelsSelected
	ErrNoAlbumsSelected
	ErrNoFilesForDownload
	ErrZipFailed

	MsgChangesSaved
	MsgAlbumCreated
	MsgAlbumSaved
	MsgAlbumDeleted
	MsgAlbumCloned
	MsgFileUnstacked
	MsgSelectionAddedTo
	MsgEntryAddedTo
	MsgEntriesAddedTo
	MsgEntryRemovedFrom
	MsgEntriesRemovedFrom
	MsgAccountCreated
	MsgAccountSaved
	MsgAccountDeleted
	MsgSettingsSaved
	MsgPasswordChanged
	MsgImportCompletedIn
	MsgImportCanceled
	MsgIndexingCompletedIn
	MsgIndexingOriginals
	MsgIndexingFiles
	MsgIndexingCanceled
	MsgRemovedFilesAndPhotos
	MsgMovingFilesFrom
	MsgCopyingFilesFrom
	MsgLabelsDeleted
	MsgLabelSaved
	MsgFilesUploadedIn
	MsgSelectionArchived
	MsgSelectionRestored
	MsgSelectionProtected
	MsgAlbumsDeleted
	MsgZipCreatedIn
)

var MsgEnglish = MessageMap{
	// Error messages:
	ErrUnexpected:         "Unexpected error, please try again",
	ErrBadRequest:         "Invalid request",
	ErrSaveFailed:         "Changes could not be saved",
	ErrDeleteFailed:       "Could not be deleted",
	ErrAlreadyExists:      "%s already exists",
	ErrNotFound:           "Not found on server, deleted?",
	ErrFileNotFound:       "File not found",
	ErrSelectionNotFound:  "Selection not found",
	ErrEntityNotFound:     "Not found on server, deleted?",
	ErrAccountNotFound:    "Account not found",
	ErrUserNotFound:       "User not found",
	ErrLabelNotFound:      "Label not found",
	ErrAlbumNotFound:      "Album not found",
	ErrPublic:             "Not available in public mode",
	ErrReadOnly:           "not available in read-only mode",
	ErrUnauthorized:       "Please log in and try again",
	ErrOffensiveUpload:    "Upload might be offensive",
	ErrNoItemsSelected:    "No items selected",
	ErrCreateFile:         "Failed creating file, please check permissions",
	ErrCreateFolder:       "Failed creating folder, please check permissions",
	ErrConnectionFailed:   "Could not connect, please try again",
	ErrInvalidPassword:    "Invalid password, please try again",
	ErrFeatureDisabled:    "Feature disabled",
	ErrNoLabelsSelected:   "No labels selected",
	ErrNoAlbumsSelected:   "No albums selected",
	ErrNoFilesForDownload: "No files available for download",
	ErrZipFailed:          "Failed to create zip file",

	// Info and confirmation messages:
	MsgChangesSaved:          "Changes successfully saved",
	MsgAlbumCreated:          "Album created",
	MsgAlbumSaved:            "Album saved",
	MsgAlbumDeleted:          "Album %s deleted",
	MsgAlbumCloned:           "Album contents cloned",
	MsgFileUnstacked:         "File removed from stack",
	MsgSelectionAddedTo:      "Selection added to %s",
	MsgEntryAddedTo:          "One entry added to %s",
	MsgEntriesAddedTo:        "%d entries added to %s",
	MsgEntryRemovedFrom:      "One entry removed from %s",
	MsgEntriesRemovedFrom:    "%d entries removed from %s",
	MsgAccountCreated:        "Account created",
	MsgAccountSaved:          "Account saved",
	MsgAccountDeleted:        "Account deleted",
	MsgSettingsSaved:         "Settings saved",
	MsgPasswordChanged:       "Password changed",
	MsgImportCompletedIn:     "Import completed in %d s",
	MsgImportCanceled:        "Import canceled",
	MsgIndexingCompletedIn:   "Indexing completed in %d s",
	MsgIndexingOriginals:     "Indexing originals...",
	MsgIndexingFiles:         "Indexing files in %s",
	MsgIndexingCanceled:      "Indexing canceled",
	MsgRemovedFilesAndPhotos: "Removed %d files and %d photos",
	MsgMovingFilesFrom:       "Moving files from %s",
	MsgCopyingFilesFrom:      "Copying files from %s",
	MsgLabelsDeleted:         "Labels deleted",
	MsgLabelSaved:            "Label saved",
	MsgFilesUploadedIn:       "%d files uploaded in %d s",
	MsgSelectionArchived:     "Selection archived",
	MsgSelectionRestored:     "Selection restored",
	MsgSelectionProtected:    "Selection marked as private",
	MsgAlbumsDeleted:         "Albums deleted",
	MsgZipCreatedIn:          "Zip created in %d s",
}
