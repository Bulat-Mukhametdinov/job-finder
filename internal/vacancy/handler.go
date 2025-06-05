package vacancy

import (
	"encoding/json"
	"html/template"
	"job-finder/internal/app"
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

type JobHandler struct {
	*app.App
}

func NewJobHandler(app *app.App) *JobHandler {
	return &JobHandler{app}
}

func (h *JobHandler) BasePage(w http.ResponseWriter, r *http.Request) {
	if len(r.RequestURI) >= 2 && r.RequestURI[1] != '?' {
		w.WriteHeader(404)
		return
	}

	user, _ := middleware.GetUserFromContext(r.Context())

	query := r.URL.Query().Get("q")

	jobs, err := h.Rapid.GetJobs(query, "", "", "", "", "", "", "", "", "")

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if user != nil {
		favouriteJobs, err := h.FavouriteStorage.GetByUserID(user.ID)
		if err != nil {
			log.Println("Error on retrieving favourite jobs:", err)
		} else {
			favMap := make(map[string]bool, len(favouriteJobs))
			for _, fav := range favouriteJobs {
				favMap[fav.ID] = true
			}

			for i := range jobs {
				if favMap[jobs[i].JobID] {
					jobs[i].IsFavourite = true
				}
			}
		}
	}

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

func (h *JobHandler) Favourite(w http.ResponseWriter, r *http.Request) {
	type FavouriteRequest struct {
		VacancyID string `json:"vacancyId"`
		Success   bool   `json:"success"`
	}

	w.Header().Set("Content-Type", "application/json")
	var favReq FavouriteRequest

	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		json.NewEncoder(w).Encode(favReq)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&favReq); err != nil && favReq.VacancyID != "" {
		json.NewEncoder(w).Encode(favReq)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, _ := middleware.GetUserFromContext(r.Context())

	if user == nil {
		json.NewEncoder(w).Encode(favReq)
		http.Error(w, "Method allowed only for authorized users", http.StatusUnauthorized)
		return
	}

	favourite := models.Favourite{ID: favReq.VacancyID, UserID: user.ID}

	var err error

	if r.Method == http.MethodPost {
		err = h.FavouriteStorage.Create(&favourite)
	} else if r.Method == http.MethodDelete {
		err = h.FavouriteStorage.Delete(&favourite)
	}

	if err != nil {
		log.Printf("Error on %v favourite vacancy: %v", r.Method, err)
		json.NewEncoder(w).Encode(favReq)
		http.Error(w, "Some error occured", http.StatusInternalServerError)
		return
	}

	favReq.Success = true

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(favReq)
}
