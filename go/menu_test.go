package openapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetMenuItemsHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/testmenuitems", GetTestMenuItems)
	req, _ := http.NewRequest("GET", "/testmenuitems", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var menuitems []MenuItem
	json.Unmarshal(w.Body.Bytes(), &menuitems)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, menuitems)
}
