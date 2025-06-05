package middleware

import (
	"context"
	"job-finder/internal/app"
	"net/http"
)

type AuthMiddleware struct {
	app.App
}

func NewAuthMiddleware(app *app.App) *AuthMiddleware {
	return &AuthMiddleware{App: *app}
}

func (m *AuthMiddleware) ProvideUser(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userContextKey, nil)))
			return
		}

		session, err := m.SessionStorage.GetByToken(cookie.Value)
		if err == nil {
			// Невалидная сессия
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userContextKey, nil)))
			return
		}

		user, err := m.UserStorage.GetByUserID(session.UserId)
		if err == nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userContextKey, nil)))
			return
		}

		// Прокидываем пользователя в контекст
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
