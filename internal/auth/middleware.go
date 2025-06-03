package auth

import (
	"fmt"
	"net/http"
	"job-finder/internal/storage"
	"context"
	"time"
	"job-finder/internal/auth/utils"
)

type AuthMiddleware struct{
	SessionStorage *storage.SessionStorage
}

func NewAuthMiddleware(sessionStorage *storage.SessionStorage) *AuthMiddleware{
	return &AuthMiddleware{SessionStorage: sessionStorage}
}

func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		token := utils.GetSessionToken(r)
		if token == ""{
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return 
		}

		session, err := m.SessionStorage.GetByToken(token)
		if err != nil{
			utils.ClearSessionCookie(w)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", session.UserId)
        r = r.WithContext(ctx)

		next(w, r)
	}
}

func GetUserID(r *http.Request) int {
    if userID, ok := r.Context().Value("user_id").(int); ok {
        return userID
    }
    return 0
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    userID := GetUserID(r)
    if userID == 0 {
        http.Error(w, "No user found", http.StatusInternalServerError)
        return
    }
    
    fmt.Fprintf(w, "Profile page for user ID: %d", userID)
}

func (m *AuthMiddleware) RedirectIfAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := utils.GetSessionToken(r)
		if token != "" {
			session, err := m.SessionStorage.GetByToken(token)
			if err == nil && session != nil && session.ExpiresAt.After(time.Now()) {
				// Пользователь уже залогинен — редиректим на /profile
				http.Redirect(w, r, "/profile", http.StatusFound)
				return
			}
		}
		// Гость — пропускаем дальше
		next(w, r)
	}
}




