package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NoNamePL/semaphore-demo-go-gin/internal/test"
)

func TestShowIndexPageUnAuthentication(t *testing.T) {
	r := test.GetRouter(true)
	r.GET("/", ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	test.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http  status code is 200
		statusOk := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more details tests using libraries that can
		// parse and process HTML pages
		p, err := io.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOk && pageOk
	})

}

func TestArticleUnAuthentication(t *testing.T) {
	
}