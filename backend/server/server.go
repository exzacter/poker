package server

import (
	"log"
	"net/http"
	"poker/backend/handlers"
)

func StartServer() {
	// Serve pages on the site
	http.HandleFunc("/", handlers.serveHome)
	http.HandleFunc("/register", handlers.serveRegister)
	http.HandleFunc("/login", handlers.serveLogin)

	// Handle form submissions for registration and login
	http.HandleFunc("/register-user", handlers.registerUser)
	http.HandleFunc("/login-user", handlers.loginUser)

	// Start the server
	log.Fatal(http.ListenAndServe(":8008", nil))
}

