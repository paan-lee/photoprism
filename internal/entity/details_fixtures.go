package entity

type DetailsMap map[string]Details

func (m DetailsMap) Get(name string, photoId uint) Details {
	if result, ok := m[name]; ok {
		result.PhotoID = photoId
		return result
	}

	return Details{PhotoID: photoId}
}

func (m DetailsMap) Pointer(name string, photoId uint) *Details {
	if result, ok := m[name]; ok {
		result.PhotoID = photoId
		return &result
	}

	return &Details{PhotoID: photoId}
}

var DetailsFixtures = DetailsMap{
	"lake": {
		PhotoID:   1000000,
		Keywords:  "nature, frog",
		Notes:     "notes",
		Subject:   "Lake",
		Artist:    "Hans",
		Copyright: "copy",
		License:   "MIT",
	},
	"blacklist": {
		PhotoID:   1000001,
		Keywords:  "screenshot, info",
		Notes:     "notes",
		Subject:   "Blacklist",
		Artist:    "Hans",
		Copyright: "copy",
		License:   "MIT",
	},
}
