package gomax

import (
	"fmt"
	"net/http"
	"log"
)

type App struct {
	router *Router
}

func NewApp() *App {
	return &App{}
}

func (a *App) SetRouter(router *Router) {
	a.router = router
}

func (a *App) Init() {
	fmt.Println("Initialize app")
}

func (a *App) Run() {
	log.Printf("Gomax version %s\n", VERSION)
	fmt.Println("Starting development server at http://127.0.0.1:3000/")
	fmt.Println("Quit the server with Ctrl + C")
	log.Fatal(http.ListenAndServe("localhost:3000", a.router))
}
