package middleware

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDefaultRecorded(t *testing.T) {
	r := gin.New()
	r.Use(New(Options{}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	if w.Code != http.StatusOK {
		t.Error("Incorrect HTTP status")
	}
	if string(responseData) != "{\"message\":\"pong\"}" {
		t.Error("Incorrect HTTP body")
	}
}
