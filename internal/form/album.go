package form

import "github.com/ulule/deepcopier"

// Album represents an album edit form.
type Album struct {
	CoverUID         string `json:"CoverUID"`
	FolderUID        string `json:"FolderUID"`
	AlbumType        string `json:"Type"`
	AlbumTitle       string `json:"Title"`
	AlbumCategory    string `json:"Category"`
	AlbumCaption     string `json:"Caption"`
	AlbumDescription string `json:"Description"`
	AlbumNotes       string `json:"Notes"`
	AlbumFilter      string `json:"Filter"`
	AlbumOrder       string `json:"Order"`
	AlbumTemplate    string `json:"Template"`
	AlbumCountry     string `json:"Country"`
	AlbumYear        int    `json:"Year"`
	AlbumMonth       int    `json:"Month"`
	AlbumFavorite    bool   `json:"Favorite"`
	AlbumPrivate     bool   `json:"Private"`
}

func NewAlbum(m interface{}) (f Album, err error) {
	err = deepcopier.Copy(m).To(&f)

	return f, err
}
