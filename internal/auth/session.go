package auth

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("super_secret_key"))