package main

import (
	"backend/internal/models"
	"encoding/json"
	"errors"
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

	writeJSONResponse(w, http.StatusOK, movies)
}

func (s *WebServer) Authenticate(w http.ResponseWriter, r *http.Request) {
	// read JSON payload
	var requestPayload models.UserAuthPayload

	w.Header().Set("Access-Control-Allow-Credentials", "true")

	err := readJSONPayload(w, r, &requestPayload)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := s.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		writeJSONError(w, errors.New("invalid credentials"), http.StatusNotFound)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password) // compare against hash
	if err != nil || !valid {
		writeJSONError(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	jwtUser := JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := s.Auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	refreshCookie := s.Auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	writeJSONResponse(w, http.StatusAccepted, tokens)
}
