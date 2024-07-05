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
	Domain string
	Port   string
	DSN    string
	DB     respository.DatabaseRepo
	Auth   Auth
}

func main() {
	// read environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	var server WebServer

	// set application config
	server.Port = os.Getenv("PORT")
	server.Domain = os.Getenv("DOMAIN")

	server.Auth = Auth{
		TokenExpiry:   15 * time.Minute,
		RefreshExpiry: 24 * time.Hour,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  server.Domain,
	}

	defaultDsn := os.Getenv("DEFAULT_DSN")

	// read from command line
	flag.StringVar(&server.DSN, "dsn", defaultDsn, "postgres connection string")
	flag.StringVar(&server.Auth.Secret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&server.Auth.Issuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&server.Auth.Audience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&server.Auth.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.Parse()

	// connect to the database
	conn, err := server.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	server.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer server.DB.Connection().Close()

	// start a web server
	log.Println("server running on port", server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, server.Routes()))
}
