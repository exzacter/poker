package handlers

import (
	"fmt"
	"net/http"
	"poker/backend/db"
	"poker/backend/models"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/templates/home.html")
}

func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/templates/login.html")
}

func serveRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/templates/register.html")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	confPassword := r.FormValue("conf-password")
	dob := r.FormValue("dob")
	email := r.FormValue("email")

	// Check if passwords match
	if password != confPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Insert user into database
	_, err := db.DB.Exec("INSERT INTO users (username, password, dob, email) VALUES ($1, $2, $3, $4)", username, password, dob, email)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %s registered successfully", username)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user models.User
	err := db.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password)

	if err != nil || user.Password != password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "User %s logged in successfully!", username)
}
	
