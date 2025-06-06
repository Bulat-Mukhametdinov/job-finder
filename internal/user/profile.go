package user

import (
	"html/template"
	"job-finder/internal/app"
	"job-finder/internal/client/rapid"
	"job-finder/internal/middleware"
	"job-finder/internal/models"
	"log"
	"net/http"
	"path/filepath"
)

type ProfileHandler struct {
	*app.App
}

func NewProfileHandler(app *app.App) *ProfileHandler {
	return &ProfileHandler{app}
}

type TemplateData struct {
	User *models.User
	Jobs []rapid.Job
}

func (h *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	favouriteJobs, err := h.FavouriteStorage.GetByUserID(user.ID)
	jobs := make([]rapid.Job, 0, len(favouriteJobs))

	if err != nil {
		log.Println("Error on retrieving favourite jobs in profile:", err)

	} else {
		for _, fav_job := range favouriteJobs {
			job, err := h.Rapid.GetJob(fav_job.ID)
			if err != nil {
				continue
			}
			job.IsFavourite = true
			job.JobComment = fav_job.Comments.String
			jobs = append(jobs, job)
		}
	}

	data := TemplateData{
		User: user,
		Jobs: jobs,
	}

	tmpl, err := template.ParseFiles(filepath.Join("web", "templates", "profile.html"))
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
