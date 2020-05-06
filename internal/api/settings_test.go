package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"net/http"
	"testing"
)

func TestGetSettings(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, router, conf := NewApiTest()
		GetSettings(router, conf)
		r := PerformRequest(app, "GET", "/api/v1/settings")
		val := gjson.Get(r.Body.String(), "theme")
		assert.NotEmpty(t, val.String())
		val2 := gjson.Get(r.Body.String(), "language")
		assert.NotEmpty(t, val2.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestSaveSettings(t *testing.T) {
	/* t.Run("successful request", func(t *testing.T) {
		app, router, conf := NewApiTest()
		GetSettings(router, conf)
		r := PerformRequest(app, "GET", "/api/v1/settings")
		val := gjson.Get(r.Body.String(), "language")
		assert.Equal(t, "de", val.String())
		assert.Equal(t, http.StatusOK, r.Code)

		SaveSettings(router, conf)
		r2 := PerformRequestWithBody(app, "POST", "/api/v1/settings", `{"language": "en"}`)
		val2 := gjson.Get(r2.Body.String(), "language")
		assert.Equal(t, "en", val2.String())
		assert.Equal(t, http.StatusOK, r2.Code)
		r3 := PerformRequestWithBody(app, "POST", "/api/v1/settings", `{"language": "de"}`)
		assert.Equal(t, http.StatusOK, r3.Code)
	}) */
	t.Run("bad request", func(t *testing.T) {
		app, router, conf := NewApiTest()
		SaveSettings(router, conf)
		r := PerformRequestWithBody(app, "POST", "/api/v1/settings", `{"language": 123}`)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}
