package main

import (
	"fmt"
	"net/http"

	"gomax/app"
)

func main() {
	a := app.App{}
	a.Init()
	a.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!!!")
	})

	http.ListenAndServe(":3000", nil)
}
