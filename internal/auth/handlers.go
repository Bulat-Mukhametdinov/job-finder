package auth

import ("net/http"
		"time"
		"job-finder/internal/storage"
		"job-finder/internal/models"
		"github.com/jmoiron/sqlx"
	)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB){
	if r.Method != http.MethodPost{
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashed := HashPassword(password)
	user := models.User{Username: username, PasswordHash: hashed}
	userStorage := &storage.UserStorage{DB:db}
	if err := userStorage.Create(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method")
	}
}

