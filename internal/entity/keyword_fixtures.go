package entity

type KeywordMap map[string]Keyword

var KeywordFixtures = KeywordMap{
	"bridge": {
		ID:      1000000,
		Keyword: "bridge",
		Skip:    false,
	},
	"beach": {
		ID:      1000001,
		Keyword: "beach",
		Skip:    false,
	},
}

// CreateKeywordFixtures inserts known entities into the database for testing.
func CreateKeywordFixtures() {
	for _, entity := range KeywordFixtures {
		Db().Create(&entity)
	}
}
