package main

import (
	"fmt"
	"job-finder/internal/app"
	"job-finder/internal/middleware"
	"job-finder/internal/storage"
	"job-finder/internal/user"
	"job-finder/internal/vacancy"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := storage.ConnectSQLite("data.db")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	err = storage.RunMigrations(db)
	if err != nil {
		log.Fatal("Error on migrations:", err)
	}

	app := app.NewApp(db)
	mdlw := middleware.NewAuthMiddleware(app)
	mux := http.NewServeMux()

	mux.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))
	user.RegisterRoutes(mux, app, mdlw)
	vacancy.RegisterRoutes(mux, app, mdlw)

	loggedMux := middleware.Logger(mux)

	fmt.Println("Server is up on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)

}
