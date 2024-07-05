package main

import (
	"backend/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
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
	jwtUser := JwtUser{ID: user.ID,
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

func (s *WebServer) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	for _, cookie := range r.Cookies() {
		fmt.Println(cookie.Name, s.Auth.CookieName)
		if cookie.Name == s.Auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (any, error) {
				return []byte(s.Auth.Secret), nil
			})
			if err != nil {
				writeJSONError(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get user id from the token claims
			userId, err := strconv.Atoi(claims.Subject) // subject is user id
			if err != nil {
				writeJSONError(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := s.DB.GetUserById(userId)
			if err != nil {
				writeJSONError(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			jwtUser := JwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokens, err := s.Auth.GenerateTokenPair(&jwtUser)
			if err != nil {
				writeJSONError(w, errors.New("error generating tokens"), http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, s.Auth.GetRefreshCookie(tokens.RefreshToken))
			writeJSONResponse(w, http.StatusOK, tokens)
		}
	}
}

func (s *WebServer) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	http.SetCookie(w, s.Auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}
