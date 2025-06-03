package main

import (
	"fmt"
	"job-finder/internal/handler"
	"job-finder/internal/storage"
	"log"
	"net/http"
	"job-finder/internal/auth"
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
	http.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))

	http.HandleFunc("/", handler.JobHandler)

	fmt.Println("Server is up on http://localhost:8080")
	auth.RegisterRoutes(db)
	http.ListenAndServe(":8080", nil)
	

}
