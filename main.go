package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// sample data

type userDetails struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []userDetails{
	{ID: 1, Name: "dhinesh", Email: "dhinesh@samplemail.com", Password: "kjshdkfhskdfh"},
	{ID: 2, Name: "dhinesh1", Email: "dhinesh1@samplemail.com", Password: "kjshdkfhskdfh"},
}

func getUsersData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func main() {
	log.Print("starting server...")
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/post", postHandler)

	// get users data
	http.HandleFunc("/users", getUsersData)

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

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

// function for post handler
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post server")
}
