package config

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestTestCliContext(t *testing.T) {
	result := CliTestContext()

	assert.IsType(t, new(cli.Context), result)
}

func TestTestConfig(t *testing.T) {
	c := TestConfig()

	assert.IsType(t, new(Config), c)
	assert.IsType(t, &gorm.DB{}, c.Db())
}

func TestNewTestParams(t *testing.T) {
	c := NewTestParams()

	assert.IsType(t, new(Params), c)

	assert.Equal(t, fs.Abs("../../assets"), c.AssetsPath)
	assert.True(t, c.Debug)
}

func TestNewTestParamsError(t *testing.T) {
	c := NewTestParamsError()

	assert.IsType(t, new(Params), c)

	assert.Equal(t, fs.Abs("../.."), c.AssetsPath)
	assert.Equal(t, "../../assets/testdata/cache", c.CachePath)
	assert.False(t, c.Debug)
}

func TestNewTestErrorConfig(t *testing.T) {
	c := NewTestErrorConfig()

	db := c.Db()

	assert.IsType(t, &gorm.DB{}, db)
}
