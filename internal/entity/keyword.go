package entity

import (
	"strings"

	"github.com/photoprism/photoprism/pkg/txt"
)

// Keyword used for full text search
type Keyword struct {
	ID      uint   `gorm:"primary_key"`
	Keyword string `gorm:"type:varchar(64);index;"`
	Skip    bool
}

// NewKeyword registers a new keyword in database
func NewKeyword(keyword string) *Keyword {
	keyword = strings.ToLower(txt.Clip(keyword, txt.ClipKeyword))

	result := &Keyword{
		Keyword: keyword,
	}

	return result
}

// Updates multiple columns in the database.
func (m *Keyword) Updates(values interface{}) error {
	return UnscopedDb().Model(m).UpdateColumns(values).Error
}

// Updates a column in the database.
func (m *Keyword) Update(attr string, value interface{}) error {
	return UnscopedDb().Model(m).UpdateColumn(attr, value).Error
}

// Save updates the existing or inserts a new row.
func (m *Keyword) Save() error {
	return Db().Save(m).Error
}

// Create inserts a new row to the database.
func (m *Keyword) Create() error {
	return Db().Create(m).Error
}

// FirstOrCreateKeyword returns the existing row, inserts a new row or nil in case of errors.
func FirstOrCreateKeyword(m *Keyword) *Keyword {
	result := Keyword{}

	if err := Db().Where("keyword = ?", m.Keyword).First(&result).Error; err == nil {
		return &result
	} else if err := m.Create(); err != nil {
		log.Errorf("keyword: %s", err)
		return nil
	}

	return m
}
