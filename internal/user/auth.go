package user

import (
	"job-finder/internal/app"
	"job-finder/internal/models"
	"job-finder/internal/user/utils"
	"log"
	"net/http"
)

type AuthHandler struct {
	*app.App
}

func NewAuthHandler(app *app.App) *AuthHandler {
	return &AuthHandler{app}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token := utils.GetSessionToken(r)
	if token != "" {
		h.SessionStorage.DeleteByToken(token)
	}

	utils.ClearSessionCookie(w)
	http.Redirect(w, r, "/", http.StatusFound)
}


// Показывает страницу регистрации
func (h *AuthHandler) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Регистрация</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f0f8ff;
					color: #003366;
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
				form {
					background-color: white;
					padding: 2em;
					border-radius: 10px;
					box-shadow: 0 0 10px rgba(0, 0, 128, 0.2);
					width: 300px;
				}
				h1 {
					text-align: center;
				}
				label {
					display: block;
					margin-top: 1em;
				}
				input {
					width: 100%;
					padding: 0.5em;
					margin-top: 0.3em;
					border: 1px solid #ccc;
					border-radius: 5px;
				}
				button {
					margin-top: 1.5em;
					width: 100%;
					padding: 0.7em;
					background-color: #0066cc;
					color: white;
					border: none;
					border-radius: 5px;
					cursor: pointer;
				}
				button:hover {
					background-color: #004a99;
				}
			</style>
		</head>
		<body>
			<form method="POST" action="/register">
				<h1>Регистрация</h1>
				<label for="username">Имя пользователя:</label>
				<input type="text" id="username" name="username" required>

				<label for="password">Пароль:</label>
				<input type="password" id="password" name="password" required>

				<label for="confirm_password">Повторите пароль:</label>
				<input type="password" id="confirm_password" name="confirm_password" required>

				<button type="submit">Зарегистрироваться</button>
			
			</form>
				<script>
				document.querySelector("form").addEventListener("submit", async function(e) {
					const username = document.getElementById("username").value;
					const password = document.getElementById("password").value;
					const confirm = document.getElementById("confirm_password").value;

					if (password !== confirm) {
						e.preventDefault();
						alert("Пароли не совпадают!");
						return;
					}

					// Проверка существующего имени
					const res = await fetch("/check-username?username=" + encodeURIComponent(username));
					const data = await res.json();
					if (data.exists) {
						e.preventDefault();
						alert("Имя пользователя уже занято!");
					}
				});
				</script>

		</body>
		</html>
	`))
}


// Обрабатывает POST-запрос на регистрацию (ваш существующий Register)
func (h *AuthHandler) ProcessRegistration(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if password != confirmPassword {
		http.Error(w, "Пароли не совпадают", http.StatusBadRequest)
		return
	}

	hashed := HashPassword(password)
	user := models.User{
		Username:     username,
		PasswordHash: hashed,
	}

	if err := h.UserStorage.Create(&user); err != nil {
		http.Error(w, "Ошибка при создании пользователя: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем ID созданного пользователя, если Create его не возвращает — можно вызвать GetByUsername
	userInDB, err := h.UserStorage.GetByUsername(user.Username)
	if err != nil {
		http.Error(w, "Ошибка при получении пользователя после регистрации", http.StatusInternalServerError)
		return
	}

	// Удалим возможные старые сессии (на всякий случай)
	h.SessionStorage.DeleteByUserID(userInDB.ID)

	// Создаем сессию
	session := utils.CreateSession(userInDB.ID)

	if err := h.SessionStorage.Create(session); err != nil {
		http.Error(w, "Ошибка при создании сессии: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SetSessionCookie(w, session.Token)

	log.Printf("User %v successfully registered and logged in\n", username)
	http.Redirect(w, r, "/", http.StatusFound)
}


// Аналогично для Login: ShowLoginPage и ProcessLogin
func (h *AuthHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Вход</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f0f8ff;
					color: #003366;
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
				form {
					background-color: white;
					padding: 2em;
					border-radius: 10px;
					box-shadow: 0 0 10px rgba(0, 0, 128, 0.2);
					width: 300px;
				}
				h1 {
					text-align: center;
					margin-bottom: 1em;
				}
				label {
					display: block;
					margin-top: 1em;
				}
				input {
					width: 100%;
					padding: 0.5em;
					margin-top: 0.3em;
					border: 1px solid #ccc;
					border-radius: 5px;
				}
				button {
					margin-top: 1.5em;
					width: 100%;
					padding: 0.7em;
					background-color: #0066cc;
					color: white;
					border: none;
					border-radius: 5px;
					cursor: pointer;
				}
				button:hover {
					background-color: #004a99;
				}
			</style>
		</head>
		<body>
			<form method="POST" action="/login">
				<h1>Вход</h1>
				<label for="username">Имя пользователя:</label>
				<input type="text" id="username" name="username" required>

				<label for="password">Пароль:</label>
				<input type="password" id="password" name="password" required>

				<button type="submit">Войти</button>
			</form>

		</body>
		</html>
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
	http.Redirect(w, r, "/", http.StatusFound)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Login successful"))

}
