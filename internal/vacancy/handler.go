package vacancy

import (
	"html/template"
	"job-finder/internal/client/rapid"
	"job-finder/internal/middleware"
	"job-finder/internal/models"
	"log"
	"net/http"
	"path/filepath"
)

type TemplateData struct {
	User  *models.User
	Query string
	Jobs  []rapid.Job
}

func JobHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())

	query := r.URL.Query().Get("q")

	jobs := rapid.GetJob(query, "", "", "", "", "", "", "", "", "")

	data := TemplateData{
		User:  user,
		Query: query,
		Jobs:  jobs,
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
