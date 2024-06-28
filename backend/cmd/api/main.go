package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type WebServer struct {
	Domain string
	Port   string
}

func main() {
	// set application config
	var server WebServer

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	server.Port = os.Getenv("PORT")
	server.Domain = os.Getenv("DOMAIN")

	// read from command line

	// connect to the database

	http.HandleFunc("/", Hello)

	// start a web server
	log.Println("server running on port", server.Port)
	log.Fatal(http.ListenAndServe(":"+server.Port, nil))
}
