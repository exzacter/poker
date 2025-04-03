package main

import (
	"log"
	"poker/backend/db"
	"poker/backend/server"
)

func main() {
	// start my database connection
	db.InitDB()
	defer db.CloseDB()
	
	// start http server
	log.Println("Server running on :8008")
	server.StartServer()

/*
	connStr := "postgres://postgres:password@11.0.0.196:5432/gopgwebsite?sslmode=disable"

	db, err1 := sql.Open("postgres", connStr)

	defer db.Close()

	if err1 != nil {
		log.Fatal(err1)
	}

	if err1 = db.Ping(); err1 != nil {
		log.Fatal(err1)
	}

	createUser(db)
*/	
}

