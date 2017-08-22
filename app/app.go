package app

import (
	"fmt"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	fmt.Println("Initialize app")
	// a.Router = mux.NewRouter()
}

func (a *App) Run() {
	fmt.Println("Run app")
}
