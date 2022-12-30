package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"local.package/controller"
	"local.package/model"
	"local.package/router"

	"github.com/stretchr/testify/assert"
)

var m = model.NewModel()
var c = controller.NewController(m)
var r = router.NewRouter(c)

var s = r.SetupRouter()

func TestPingRouter(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	s.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
