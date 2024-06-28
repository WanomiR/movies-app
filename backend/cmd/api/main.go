package main

import (
	"backend/internal/respository"
	"backend/internal/respository/dbrepo"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Domain       string
	Port         string
	DSN          string
	DB           respository.DatabaseRepo
	Auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	// set application config
	var server WebServer

	// read environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	server.Port = os.Getenv("PORT")
	defaultDsn := os.Getenv("DEFAULT_DSN")

	// read from command line
	flag.StringVar(&server.DSN, "dsn", defaultDsn, "postgres connection string")
	flag.StringVar(&server.JWTSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&server.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&server.JWTAudience, "jwt-audience", "signing audience", "signing audience")
	flag.StringVar(&server.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&server.Domain, "domain", "example.com", "domain")
	flag.Parse()

	// connect to the database
	conn, err := server.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	server.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer server.DB.Connection().Close()

	server.Auth = Auth{
		Issuer:        server.JWTIssuer,
		Audience:      server.JWTAudience,
		Secret:        server.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		Cookie: Cookie{
			Path:   "/",
			Name:   "__Host-refresh_token",
			Domain: server.CookieDomain,
		},
	}

	// start a web server
	log.Println("server running on port", server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, server.Routes()))
}
