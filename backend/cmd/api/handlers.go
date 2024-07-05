package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *WebServer) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println(err)
	}
}

func (s *WebServer) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := s.DB.AllMovies()
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		log.Println(err)
	}
}

//func (s *WebServer) Authenticate(w http.ResponseWriter) {}
