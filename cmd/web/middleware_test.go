package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var my myHandler
	h := NoSurf(&my)

	switch h.(type) {
	case http.Handler:
	default:
		t.Error("The type is not a http.Handler")

	}
}

func TestSessionLoad(t *testing.T) {
	var my myHandler
	h := SessionLoad(&my)
	switch h.(type) {
	case http.Handler:
	default:
		t.Error("The type is not a http.Handler")

	}
}
