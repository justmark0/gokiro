package app

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/upload", a.UploadHandler()).Methods("POST")
	a.Router.HandleFunc("/download", a.DownloadHandler()).Methods("GET")
}
