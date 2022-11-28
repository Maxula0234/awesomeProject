package main

import (
	"awesomeProject/internal/handlers"
	"awesomeProject/version"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Start service")
	log.Printf(
		"Starting the service...\ncommit: %s, \nbuild time: %s, \nrelease: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	//port := "8000"

	r := handlers.Router()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
