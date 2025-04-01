package main

import (
	"log"
	"net/http"
	//"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"database/sql"
)

func main() {
	
	connStr := "postgres://postgres:password@10.0.0.196:5432/gopgwebsite?sslmode=disable"

	db, err1 := sql.Open("postgres", connStr)

	defer db.Close()

	if err1 != nil {
		log.Fatal(err1)
	}

	if err1 = db.Ping(); err1 != nil {
		log.Fatal(err1)
	}

	createUser(db)
	


	pokertable := http.NewServeMux()

	pokertable.HandleFunc("GET /{$}", home)
	pokertable.HandleFunc("GET /register", register )
	pokertable.HandleFunc("GET /login", login)
	// Print a log message to say that the server is starting
	log.Print("Starting server on :8040")

	// Use the ListenAndServe() function to start the web server. We pass in two parameters: TCP port to listen on, and the servemux we just created.
	// we use the Log.Fatal() to log the error message and exit.
	err := http.ListenAndServe(":8040", pokertable)
	log.Fatal(err)
}
/*
func createUser(db *sql.DB) {
    query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        dob DATE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}
*/
