package handler

import (
	"html/template"
	"job-finder/internal/client/rapid"
	"log"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
)

type TemplateData struct {
	Registration string
	Query        string
	Jobs         []rapid.Job
}

func JobHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	query := r.URL.Query().Get("q")
	register := r.URL.Query().Get("register")

	jobs := rapid.GetJob(query, "", "", "", "", "", "", "", "", "")

	data := TemplateData{
		Registration: register,
		Query:        query,
		Jobs:         jobs,
	}

	tmpl, err := template.ParseFiles(filepath.Join("web", "templates", "index.html"))
	if err != nil {
		log.Println("Template parsing error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Rendering error:", err)
		http.Error(w, "Page rendering error", http.StatusInternalServerError)
	}
}
