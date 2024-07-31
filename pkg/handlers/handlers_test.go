package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var testCase = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()
	for _, e := range testCase {
		if e.method == "GET" {
			response, error := ts.Client().Get(ts.URL + e.url)
			if error != nil {
				t.Log(error)
				t.Fatal(error)
			}
			if response.StatusCode != e.expectedStatusCode {
				t.Errorf("%s expected status code %d, got %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		} else {

		}
	}
}
