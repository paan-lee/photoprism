package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhotoUnstack(t *testing.T) {
	t.Run("unstack xmp sidecar file", func(t *testing.T) {
		app, router, _ := NewApiTest()
		PhotoUnstack(router)
		r := PerformRequest(app, "POST", "/api/v1/photos/pt9jtdre2lvl0yh7/files/ft1es39w45bnlqdw/unstack")
		// Sidecar files can not be unstacked.
		assert.Equal(t, http.StatusBadRequest, r.Code)
		// t.Logf("RESP: %s", r.Body.String())
	})

	t.Run("unstack bridge3.jpg", func(t *testing.T) {
		app, router, _ := NewApiTest()
		PhotoUnstack(router)
		r := PerformRequest(app, "POST", "/api/v1/photos/pt9jtdre2lvl0yh7/files/ft2es49whhbnlqdn/unstack")
		// TODO: Have a real file in place for testing the success case. This file does not exist, so it can't be unstacked.
		assert.Equal(t, http.StatusNotFound, r.Code)
		// t.Logf("RESP: %s", r.Body.String())
	})
}
