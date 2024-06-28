package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (s *WebServer) Routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/", s.Home)
	mux.Get("/movies", s.AllMovies)

	return mux
}
