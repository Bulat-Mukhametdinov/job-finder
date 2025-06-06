```
/job-finder/
│
├── /app/                      # точка входа (main.go)
│
├── /internal/
│   ├── /app/                  # базовое приложение
│   ├── /middleware/           # middlewares
│   ├── /vacancy/              # HTTP-ручки главной страницы
│   ├── /user/                 # логика авторизации и управления пользователем
│   ├── /storage/              # доступ к БД (SQLite)
│   ├── /model/                # структуры данных
│   └── /client/rapid/         # работа с API
│
├── /web/                      # фронт
│   ├── /static/               # JS, CSS
│   └── /templates/            # HTML-шаблоны
│
├── go.mod
└── README.md
```
