package main

import (
	"log"
	"net/http"
)

func main() {
	pokertable := http.NewServeMux()

	pokertable.HandleFunc("GET /{$}", home)
	pokertable.HandleFunc("GET /login", login)
	// Print a log message to say that the server is starting
	log.Print("Starting server on :8040")

	// Use the ListenAndServe() function to start the web server. We pass in two parameters: TCP port to listen on, and the servemux we just created.
	// we use the Log.Fatal() to log the error message and exit.
	err := http.ListenAndServe(":8040", pokertable)
	log.Fatal(err)
}
