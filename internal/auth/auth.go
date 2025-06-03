package auth

import (
	"job-finder/internal/storage"
	"net/http"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(db *sqlx.DB) {
	authHandler := NewAuthHandler(db)
	authMiddleware := NewAuthMiddleware(&storage.SessionStorage{DB: db})

	http.HandleFunc("/auth/register", authMiddleware.RedirectIfAuthenticated(func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet{
			authHandler.ShowRegisterPage(w, r)
		}else if r.Method == http.MethodPost{
			authHandler.ProcessRegistration(w, r)
		}else {http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)}
	}))
	http.HandleFunc("/auth/login", authMiddleware.RedirectIfAuthenticated(func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			authHandler.ShowLoginPage(w, r)
		} else if r.Method == http.MethodPost {
			authHandler.ProcessLogin(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/auth/logout", authHandler.Logout)
	http.HandleFunc("/profile", authMiddleware.RequireAuth(profileHandler))
}
