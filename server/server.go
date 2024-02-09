package server

import (
	"alexdenkk/liseum/controllers"

	"alexdenkk/liseum/pkg/middleware"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// App - server service app struct
type App struct {
	Handler *controllers.Handler
	Server  *http.Server
	MW      *middleware.Middleware
	SignKey []byte
	DB      *gorm.DB
}

// Run - function for run service app
func (app *App) Run() error {
	app.Route()
	log.Println("server running")
	log.Println("=============")
	log.Println("by @alexdenkk")
	log.Println("=============")
	return app.Server.ListenAndServe()
}

// New - function for create new service app
func New(db *gorm.DB, key []byte, addr string) *App {
	app := &App{
		DB:      db,
		SignKey: key,
	}

	app.Server = &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.MW = middleware.New(app.SignKey)

	app.InitHandler()

	return app
}
