package api

import (
	"net/http"
	"testing"

	"github.com/tidwall/gjson"

	"github.com/stretchr/testify/assert"
)

func TestGetLabels(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		GetLabels(router)
		r := PerformRequest(app, "GET", "/api/v1/labels?count=15")
		count := gjson.Get(r.Body.String(), "#")
		assert.LessOrEqual(t, int64(4), count.Int())
		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("invalid request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		GetLabels(router)
		r := PerformRequest(app, "GET", "/api/v1/labels?xxx=15")
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestUpdateLabel(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		UpdateLabel(router)
		r := PerformRequestWithBody(app, "PUT", "/api/v1/labels/lt9k3pw1wowuy3c7", `{"Name": "Updated01", "Priority": 2}`)
		val := gjson.Get(r.Body.String(), "Name")
		assert.Equal(t, "Updated01", val.String())
		val2 := gjson.Get(r.Body.String(), "CustomSlug")
		assert.Equal(t, "updated01", val2.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("invalid request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		UpdateLabel(router)
		r := PerformRequestWithBody(app, "PUT", "/api/v1/labels/lt9k3pw1wowuy3c7", `{"Name": 123, "Priority": 4, "Uncertainty": 80}`)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("not found", func(t *testing.T) {
		app, router, _ := NewApiTest()
		UpdateLabel(router)
		r := PerformRequestWithBody(app, "PUT", "/api/v1/labels/xxx", `{"Name": "Updated01", "Priority": 4, "Uncertainty": 80}`)
		val := gjson.Get(r.Body.String(), "error")
		assert.Equal(t, "Label not found", val.String())
		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestLikeLabel(t *testing.T) {
	t.Run("like not existing label", func(t *testing.T) {
		app, router, _ := NewApiTest()
		LikeLabel(router)
		r := PerformRequest(app, "POST", "/api/v1/labels/8775789/like")
		assert.Equal(t, http.StatusNotFound, r.Code)
	})
	t.Run("like existing label", func(t *testing.T) {
		app, router, _ := NewApiTest()
		GetLabels(router)
		r2 := PerformRequest(app, "GET", "/api/v1/labels?count=1&q=likeLabel")
		t.Log(r2.Body.String())
		val := gjson.Get(r2.Body.String(), `#(Slug=="likeLabel").Favorite`)
		assert.Equal(t, "false", val.String())
		LikeLabel(router)
		r := PerformRequest(app, "POST", "/api/v1/labels/lt9k3pw1wowuy3c9/like")
		t.Log(r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)
		r3 := PerformRequest(app, "GET", "/api/v1/labels?count=1&q=likeLabel")
		t.Log(r3.Body.String())
		val2 := gjson.Get(r3.Body.String(), `#(Slug=="likeLabel").Favorite`)
		assert.Equal(t, "true", val2.String())
	})

}

func TestDislikeLabel(t *testing.T) {
	t.Run("dislike not existing label", func(t *testing.T) {
		app, router, _ := NewApiTest()

		DislikeLabel(router)

		r := PerformRequest(app, "DELETE", "/api/v1/labels/5678/like")
		assert.Equal(t, http.StatusNotFound, r.Code)
	})
	t.Run("dislike existing label", func(t *testing.T) {
		app, router, _ := NewApiTest()
		GetLabels(router)
		r2 := PerformRequest(app, "GET", "/api/v1/labels?count=3&q=landscape")
		t.Logf("HTTP BODY: %s", r2.Body.String())
		val := gjson.Get(r2.Body.String(), `#(Slug=="landscape").Favorite`)
		assert.Equal(t, "true", val.String())

		DislikeLabel(router)

		r := PerformRequest(app, "DELETE", "/api/v1/labels/lt9k3pw1wowuy3c2/like")
		assert.Equal(t, http.StatusOK, r.Code)

		r3 := PerformRequest(app, "GET", "/api/v1/labels?count=3&q=landscape")
		val2 := gjson.Get(r3.Body.String(), `#(Slug=="landscape").Favorite`)
		assert.Equal(t, "false", val2.String())
	})
}

func TestLabelThumbnail(t *testing.T) {
	t.Run("invalid type", func(t *testing.T) {
		app, router, conf := NewApiTest()
		LabelThumbnail(router)
		r := PerformRequest(app, "GET", "/api/v1/labels/lt9k3pw1wowuy3c2/t/"+conf.PreviewToken()+"/xxx")
		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("invalid label", func(t *testing.T) {
		app, router, conf := NewApiTest()
		LabelThumbnail(router)
		r := PerformRequest(app, "GET", "/api/v1/labels/xxx/t/"+conf.PreviewToken()+"/tile_500")

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("could not find original", func(t *testing.T) {
		app, router, conf := NewApiTest()
		LabelThumbnail(router)
		r := PerformRequest(app, "GET", "/api/v1/labels/lt9k3pw1wowuy3c3/t/"+conf.PreviewToken()+"/tile_500")
		assert.Equal(t, http.StatusOK, r.Code)
	})
}
