package db

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

var DB *sql.DB 

func InitDB() {
	var err error
	connstring := "postgres://postgres:password@10.0.0.196:5432/gopgwebsite?sslmode=disable"

	DB, err = sql.Open("postgres", connstring)
	if err != nil{
		log.Fatal("Failed to connect to the database:", err )
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database connection is not alive:", err)
	}

	log.Println("Database connected succesfully")
}

func CloseDB() {
	DB.Close()
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

    fmt.Println("This is working from another file") 

    _, err := db.Exec(query)
    if err != nil {
  	log.Fatal(err)
	}
}
*/
