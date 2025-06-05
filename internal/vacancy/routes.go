package vacancy

import (
	"job-finder/internal/app"
	"job-finder/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, app *app.App, mdlw *middleware.AuthMiddleware) {
	jobHandler := NewJobHandler(app)
	mux.Handle("/", mdlw.ProvideUser(http.HandlerFunc(jobHandler.BasePage)))
	mux.Handle("/api/favourites", mdlw.ProvideUser(http.HandlerFunc(jobHandler.Favourite)))
	mux.Handle("/api/comment-favourite", mdlw.ProvideUser(http.HandlerFunc(jobHandler.LeaveComment)))
}
