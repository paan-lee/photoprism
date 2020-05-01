package entity

var PhotoLabelFixtures = map[string]PhotoLabel{
	"1": {
		PhotoID:     1000000,
		LabelID:     1000001,
		LabelSrc:    "image",
		Uncertainty: 38,
		Photo:       &PhotoFixture19800101_000002_D640C559,
		Label:       &LabelFixtureFlower,
	},
	"2": {
		PhotoID:     1000000,
		LabelID:     1000002,
		LabelSrc:    "image",
		Uncertainty: 38,
		Photo:       &PhotoFixture19800101_000002_D640C559,
		Label:       &LabelFixtureCake,
	},
}

// CreatePhotoLabelFixtures inserts known entities into the database for testing.
func CreatePhotoLabelFixtures() {
	for _, entity := range PhotoLabelFixtures {
		Db().Create(&entity)
	}
}
