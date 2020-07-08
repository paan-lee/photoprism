package entity

import "time"

var date = time.Date(2050, 3, 6, 2, 6, 51, 0, time.UTC)

type LinkMap map[string]Link

var LinkFixtures = LinkMap{
	"1jxf3jfn2k": {
		LinkToken:   "1jxf3jfn2k",
		LinkExpires: 0,
		ShareUID:    "st9lxuqxpogaaba7",
		CanComment:  true,
		CanEdit:     false,
		CreatedAt:   time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC),
	},
	"4jxf3jfn2k": {
		LinkToken:   "4jxf3jfn2k",
		LinkExpires: 0,
		ShareUID:    "at9lxuqxpogaaba7",
		CanComment:  true,
		CanEdit:     false,
		CreatedAt:   time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC),
	},
}

// CreateLinkFixtures inserts known entities into the database for testing.
func CreateLinkFixtures() {
	for _, entity := range LinkFixtures {
		Db().Create(&entity)
	}
}
