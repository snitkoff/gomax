package gomax

import (
	"fmt"
	"log"
	"net/http"
)

// App represents a main application structure
type App struct {
	router *Router
}

// NewApp returns a new empty application
func NewApp() *App {
	return &App{}
}

// SetRouter using for initialize application router
func (a *App) SetRouter(router *Router) {
	a.router = router
}

// func (a *App) Init() {
// 	fmt.Println("Initialize app")
// }

// Run function using for runing development server
func (a *App) Run() {
	log.Printf("Gomax version %s\n", Version)
	fmt.Println("Starting development server at http://127.0.0.1:3000/")
	fmt.Println("Quit the server with Ctrl + C")
	log.Fatal(http.ListenAndServe("localhost:3000", a.router))
}
