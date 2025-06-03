package auth

import ("net/http"
		"job-finder/internal/storage"
		"job-finder/internal/models"
		"github.com/jmoiron/sqlx"
		"job-finder/internal/auth/utils"
	)

type AuthHandler struct {
    UserStorage    *storage.UserStorage
    SessionStorage *storage.SessionStorage
}

func NewAuthHandler(db *sqlx.DB) *AuthHandler {
    return &AuthHandler{
        UserStorage:    &storage.UserStorage{DB: db},
        SessionStorage: &storage.SessionStorage{DB: db},
    }
}


func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
    token := utils.GetSessionToken(r)
    if token != "" {
        h.SessionStorage.DeleteByToken(token)
    }
    
    utils.ClearSessionCookie(w)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Logout successful"))
}

// В вашем auth.go или handlers.go

// Показывает страницу регистрации
func (h *AuthHandler) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	// Здесь вы должны отрендерить HTML-шаблон с формой регистрации
	// Примерно так (если используете html/template):
	/*
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
	*/
    // Или просто для теста:
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<h1>Регистрация</h1>
		<form method="POST" action="/auth/register">
			<div>
				<label for="username">Имя пользователя:</label>
				<input type="text" id="username" name="username" required>
			</div>
			<div>
				<label for="password">Пароль:</label>
				<input type="password" id="password" name="password" required>
			</div>
			<button type="submit">Зарегистрироваться</button>
		</form>
	`))
}

// Обрабатывает POST-запрос на регистрацию (ваш существующий Register)
func (h *AuthHandler) ProcessRegistration(w http.ResponseWriter, r *http.Request) {
    // Ваш текущий код Register, который читает r.FormValue()
    username := r.FormValue("username")
    password := r.FormValue("password")

    // ... остальной код ...
    hashed := HashPassword(password) // Убедитесь, что HashPassword определена
    user := models.User{Username: username, PasswordHash: hashed}
    if err := h.UserStorage.Create(&user); err != nil{
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User created successfully. You can now login."))
}

// Аналогично для Login: ShowLoginPage и ProcessLogin
func (h *AuthHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<h1>Вход</h1>
		<form method="POST" action="/auth/login">
			<div>
				<label for="username">Имя пользователя:</label>
				<input type="text" id="username" name="username" required>
			</div>
			<div>
				<label for="password">Пароль:</label>
				<input type="password" id="password" name="password" required>
			</div>
			<button type="submit">Войти</button>
		</form>
	`))
}

// Ваш существующий Login переименовываем в ProcessLogin
func (h *AuthHandler) ProcessLogin(w http.ResponseWriter, r *http.Request) {
    // Ваш текущий код Login
    username := r.FormValue("username")
    password := r.FormValue("password")

    user, err := h.UserStorage.GetByUsername(username)
    if err != nil {
        http.Error(w, "Invalid credentials (user not found)", http.StatusUnauthorized)
        return
    }

    if !CheckPassword(password, user.PasswordHash) { // Убедитесь, что CheckPassword определена
        http.Error(w, "Invalid credentials (password mismatch)", http.StatusUnauthorized)
        return
    }

    h.SessionStorage.DeleteByUserID(user.ID)
    session := utils.CreateSession(user.ID)
    err = h.SessionStorage.Create(session)
    if err != nil {
        http.Error(w, "Failed to create session "+err.Error(), http.StatusInternalServerError)
        return
    }
    utils.SetSessionCookie(w, session.Token)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Login successful"))
}