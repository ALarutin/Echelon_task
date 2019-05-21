package middleware_test

import (
	"github.com/ALarutin/Echelon_task/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicHeaders(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.BasicHeaders)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	resp := httptest.NewRecorder()

	r.POST("/", func(c *gin.Context) {})
	r.ServeHTTP(resp, req)

	contentType := "application/json"
	if header:= resp.Header().Get("Content-type"); header != contentType {
		t.Errorf("Middleware get wrong header:\nwant: %v\ngot: %v",
			contentType, header)
	}
}
