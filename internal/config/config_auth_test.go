package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMode(t *testing.T) {
	c := NewConfig(CliTestContext())
	c.options.Public = true
	c.options.Demo = false
	assert.Equal(t, AuthModePublic, c.AuthMode())
	c.options.Public = false
	c.options.Demo = false
	assert.Equal(t, AuthModePasswd, c.AuthMode())
	c.options.Demo = true
	assert.Equal(t, AuthModePublic, c.AuthMode())
	c.options.AuthMode = "pass"
	assert.Equal(t, AuthModePublic, c.AuthMode())
	c.options.Demo = false
	c.options.AuthMode = "pass"
	assert.Equal(t, AuthModePasswd, c.AuthMode())
	c.options.AuthMode = "password"
	assert.Equal(t, AuthModePasswd, c.AuthMode())
}

func TestAuth(t *testing.T) {
	c := NewConfig(CliTestContext())
	c.options.Public = true
	c.options.Demo = false
	assert.False(t, c.Auth())
	c.options.Public = false
	c.options.Demo = false
	assert.True(t, c.Auth())
	c.options.Demo = true
	assert.False(t, c.Auth())
}

func TestUtils_CheckPassword(t *testing.T) {
	c := NewConfig(CliTestContext())

	formPassword := "photoprism"

	c.options.AdminPassword = "$2b$10$cRhWIleqJkbaFWhBMp54VOI25RvVubxOooCWzWgdrvl5COFxaBnAy"
	check := c.CheckPassword(formPassword)
	assert.True(t, check)

	c.options.AdminPassword = "photoprism"
	check = c.CheckPassword(formPassword)
	assert.True(t, check)

	c.options.AdminPassword = "$2b$10$yprZEQzm/Qy7AaePXtKfkem0kANBZgRwl8HbLE4JrjK6/8Pypgi1W"
	check = c.CheckPassword(formPassword)
	assert.False(t, check)

	c.options.AdminPassword = "admin"
	check = c.CheckPassword(formPassword)
	assert.False(t, check)
}

func TestUtils_isBcrypt(t *testing.T) {
	p := "$2b$10$cRhWIleqJkbaFWhBMp54VOI25RvVubxOooCWzWgdrvl5COFxaBnAy"
	assert.True(t, isBcrypt(p))

	p = "$2b$10$cRhWIleqJkbaFWhBMp54VOI25RvVubxOooCWzWgdrvl5COFxaBnA"
	assert.False(t, isBcrypt(p))

	p = "admin"
	assert.False(t, isBcrypt(p))

	p = ""
	assert.False(t, isBcrypt(p))
}

func TestConfig_InvalidDownloadToken(t *testing.T) {
	c := NewConfig(CliTestContext())

	assert.True(t, c.InvalidDownloadToken("xxx"))
}

func TestConfig_InvalidPreviewToken(t *testing.T) {
	c := NewConfig(CliTestContext())

	assert.True(t, c.InvalidPreviewToken("xxx"))
}
