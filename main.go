package main

import (
	"bitsize-backend/config"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	serverPort := config.AppConfig.GetPort()
	// Set up HTTP handlers
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/", helloHandler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = serverPort
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to our bits learn application")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello user , welcome to our app !")
}
