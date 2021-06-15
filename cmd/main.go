package main

import (
	"log"
	"net/http"

	"github.com/feature_toggle/pkg/server"
	_ "github.com/lib/pq"
)

func main() {
	log.Printf("Starting server...")
	s, err := server.NewServer()

	if err != nil {
		log.Fatalf("Error while instantiating server: %s", err)
	}

	log.Printf("Server is up!")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
