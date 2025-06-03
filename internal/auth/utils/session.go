package utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
	"job-finder/internal/models"
)

func GenerateSessionToken() string {
    bytes := make([]byte, 16)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}

func CreateSession(userID int) *models.Session {
    return &models.Session{
        Token:     GenerateSessionToken(),
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(24 * time.Hour), 
		UserId:    userID,
    }
}

func SetSessionCookie(w http.ResponseWriter, token string) {
    cookie := &http.Cookie{
        Name:     "session_token",
        Value:    token,
        Path:     "/",
        MaxAge:   86400, 
        HttpOnly: true,
        Secure:   false, 
        SameSite: http.SameSiteLaxMode,
    }
    http.SetCookie(w, cookie)
}

func GetSessionToken(r *http.Request) string {
    cookie, err := r.Cookie("session_token")
    if err != nil {
        return ""
    }
    return cookie.Value
}

func ClearSessionCookie(w http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:     "session_token",
        Value:    "",
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
    }
    http.SetCookie(w, cookie)
}