package main

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
)

type Cookie struct {
	Domain, Path, Name string
}

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	Cookie        Cookie
}

type JwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokensPair(user *JwtUser) (TokenPairs, error) {
	// create a token
	token := jwt.New(jwt.SigningMethodHS256)

	// set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.FirstName + " " + user.LastName
	claims["sub"] = strconv.Itoa(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix() // issued at
	claims["typ"] = "JWT"                   // type

	// set the expiry for jwt
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	// create signed token
	signedToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// create a refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	// set the claims
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = strconv.Itoa(user.ID)
	refreshClaims["iat"] = time.Now().UTC().Unix()

	// set the expiry for refresh token
	refreshClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	// create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// create token pairs and populate with signed tokens
	tokenPairs := TokenPairs{
		Token:        signedToken,
		RefreshToken: signedRefreshToken,
	}

	// return token pairs
	return tokenPairs, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.Cookie.Name,
		Path:     j.Cookie.Path,
		Domain:   j.Cookie.Domain,
		Value:    refreshToken,
		Expires:  time.Now().UTC().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     j.Cookie.Name,
		Path:     j.Cookie.Path,
		Domain:   j.Cookie.Domain,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Secure:   true,
	}
}
