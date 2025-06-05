package user

import (
	"job-finder/internal/app"
	"job-finder/internal/middleware"
	"net/http"
)

type ProfileHandler struct {
	*app.App
}

func NewProfileHandler(app *app.App) *ProfileHandler {
	return &ProfileHandler{app}
}

func (h *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromContext(r.Context())
	if user != nil {
		w.Write([]byte(user.Username))
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
