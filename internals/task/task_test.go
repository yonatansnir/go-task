package task

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllTask(t *testing.T) {
	router := gin.Default()
	CreateTaskRoutes(router);

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/task", nil)
	router.ServeHTTP(recorder, req)

	t.Run("Test status code 200", func (t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200 got :", recorder.Code)
		}
	})
}