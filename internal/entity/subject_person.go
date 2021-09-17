package entity

import (
	"encoding/json"

	"github.com/photoprism/photoprism/pkg/txt"
)

const (
	SubjPerson = "person"
)

// People represents a list of people.
type People []Person

// Person represents a subject with type person.
type Person struct {
	SubjUID      string `json:"UID"`
	SubjName     string `json:"Name"`
	SubjAlias    string `json:"Alias"`
	SubjFavorite bool   `json:"Favorite"`
	Thumb        string `json:"Thumb"`
}

// NewPerson returns a new entity.
func NewPerson(subj Subject) *Person {
	result := &Person{
		SubjUID:      subj.SubjUID,
		SubjName:     subj.SubjName,
		SubjAlias:    subj.SubjAlias,
		SubjFavorite: subj.SubjFavorite,
		Thumb:        subj.Thumb,
	}

	return result
}

// MarshalJSON returns the JSON encoding.
func (m *Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UID      string
		Name     string
		Keywords []string `json:",omitempty"`
		Favorite bool     `json:",omitempty"`
		Thumb    string   `json:",omitempty"`
	}{
		UID:      m.SubjUID,
		Name:     m.SubjName,
		Keywords: txt.NameKeywords(m.SubjName, m.SubjAlias),
		Favorite: m.SubjFavorite,
		Thumb:    m.Thumb,
	})
}
