package vacancy

import (
	"job-finder/internal/app"
	"job-finder/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, app *app.App, mdlw *middleware.AuthMiddleware) {
	mux.Handle("/", mdlw.ProvideUser(http.HandlerFunc(JobHandler)))
}
