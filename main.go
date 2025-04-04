package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("bitsSize starting server...")
	http.HandleFunc("/welcome", welcomeHanduler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func welcomeHanduler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome our bits learn application")
}
