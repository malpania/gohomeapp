package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/malpania/beerproj/pkg/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch mux.(type) {

	case *chi.Mux:
	default:
		t.Errorf("want chi.Mux, got %T", mux)

	}
}
