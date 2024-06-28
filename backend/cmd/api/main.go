package main

import (
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type WebServer struct {
	Domain string
	Port   string
	DSN    string
	DB     *sql.DB
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
	server.Domain = os.Getenv("DOMAIN")
	defaultDsn := os.Getenv("DEFAULT_DSN")

	// read from command line
	flag.StringVar(&server.DSN, "dsn", defaultDsn, "postgres connection string")
	flag.Parse()

	// connect to the database
	conn, err := server.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	server.DB = conn

	// start a web server
	log.Println("server running on port", server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, server.Routes()))
}
