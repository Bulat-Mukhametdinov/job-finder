package main

import (
	"fmt"
	"job-finder/internal/storage"
	"log"
	"net/http"
)

func main() {
	db, err := storage.ConnectSQLite("data.db")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	err = storage.RunMigrations(db)
	if err != nil {
		log.Fatal("Error on migrations:", err)
	}

	fmt.Println("Server is up on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
