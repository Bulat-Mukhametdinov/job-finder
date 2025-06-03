/job-finder/
│
├── /app/                  # точка входа (main.go)
│
├── /internal/
│   ├── /handler/              # HTTP-ручки: login, search, favorites и т.д.
│   ├── /auth/                 # логика авторизации, куки, сессии
│   ├── /storage/              # доступ к БД (SQLite)
│   ├── /service/              # бизнес-логика
│   ├── /model/                # структуры данных
│   └── /client/hh/            # работа с API hh.ru
│
├── /web/                      # фронт (HTML + CSS + JS)
│   ├── /static/               # JS, CSS
│   └── /templates/            # HTML-шаблоны
│
├── go.mod
└── README.md
