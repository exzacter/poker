package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
)

func createUser(db *sql.DB) {
    query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        dob DATE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    fmt.Println("This is working from another file") 

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}
