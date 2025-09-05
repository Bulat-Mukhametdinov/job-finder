package user

import (
	"job-finder/internal/app"
	"job-finder/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, app *app.App, mdlw *middleware.AuthMiddleware) {
	authHandler := NewAuthHandler(app)
	profileHandler := NewProfileHandler(app)

	mux.Handle("/register", mdlw.ProvideUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.ShowRegisterPage(w, r)
		} else if r.Method == http.MethodPost {
			authHandler.ProcessRegistration(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
	mux.Handle("/login", mdlw.ProvideUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.ShowLoginPage(w, r)
		} else if r.Method == http.MethodPost {
			authHandler.ProcessLogin(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
	mux.Handle("/logout", mdlw.ProvideUser(http.HandlerFunc(authHandler.Logout)))
	mux.Handle("/profile", mdlw.ProvideUser(http.HandlerFunc(profileHandler.Profile)))
}
