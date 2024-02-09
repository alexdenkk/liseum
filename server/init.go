package server

import (
	"alexdenkk/liseum/controllers"
	"alexdenkk/liseum/db"
)

// InitHandler - function for initializing handler
func (app *App) InitHandler() {
	// repository
	repository := db.New(app.DB)
	// handler
	h := controllers.New(repository, app.SignKey)
	app.Handler = h
}
