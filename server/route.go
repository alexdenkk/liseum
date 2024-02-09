package server

import (
	"alexdenkk/liseum/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Route - function for routing
func (app *App) Route() {
	r := mux.NewRouter()

	app.RouteUsers(r)
	app.RouteClasses(r)
	app.RouteImages(r)
	app.RouteAdmin(r)

	r.Use(middleware.LoggerMW)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./web/static/")),
		),
	)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/class/all/", http.StatusMovedPermanently)
	})

	app.Server.Handler = r
}

func (app *App) RouteUsers(r *mux.Router) {
	sub := r.PathPrefix("/user").Subrouter()
	sub.HandleFunc("/login/", app.MW.NotAuth(app.Handler.LoginUser))
}

func (app *App) RouteClasses(r *mux.Router) {
	sub := r.PathPrefix("/class").Subrouter()
	sub.HandleFunc("/create/", app.MW.Auth(app.Handler.CreateClass))
	sub.HandleFunc("/all/", app.Handler.GetAllClasses)
}

func (app *App) RouteImages(r *mux.Router) {
	sub := r.PathPrefix("/img").Subrouter()
	sub.HandleFunc("/for/{class}/", app.Handler.GetImagesFor)
	sub.HandleFunc("/{class}/create/", app.MW.Auth(app.Handler.CreateImage))
}

func (app *App) RouteAdmin(r *mux.Router) {
	sub := r.PathPrefix("/admin").Subrouter()
	sub.HandleFunc("/", app.MW.Auth(app.Handler.Admin))
}
