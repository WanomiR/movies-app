package main

import (
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

	if err := s.writeJSONResponse(w, http.StatusOK, payload); err != nil {
		log.Println(err)
	}
}

func (s *WebServer) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := s.DB.AllMovies()
	if err != nil {
		s.writeJSONError(w, err)
		return
	}

	if err = s.writeJSONResponse(w, http.StatusOK, movies); err != nil {
		log.Println(err)
	}
}

func (s *WebServer) Authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")

	err := s.readJSONPayload(w, r, &requestPayload)
	if err != nil {
		s.writeJSONError(w, err)
		return
	}

	// validate user against database
	user, err := s.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		s.writeJSONError(w, errors.New("invalid credentials"))
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		s.writeJSONError(w, errors.New("invalid credentials"))
		return
	}

	// create a jwt user
	jwtUser := JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := s.Auth.GenerateTokensPair(&jwtUser)
	if err != nil {
		s.writeJSONError(w, err)
		return
	}

	// set cookie
	refreshCookie := s.Auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	s.writeJSONResponse(w, http.StatusAccepted, tokens)
}
