package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/NoNamePL/semaphore-demo-go-gin/internal/models"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

// This function is using for setup before excecuting the test functions
func TestMain(m *testing.M) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("web/templates/*")
	}
	return r
}

// Helper function to process a request and a test its response
func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the server and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is using to store the main list into the temporary one
// for testing
func SaveList() {
	tmpArticleList = models.Articlelist
}

// This function is using to restore the main list from the temporary one
func RestoreList() {
	models.Articlelist = tmpArticleList
}
