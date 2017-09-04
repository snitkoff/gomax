package gomax

import "net/http"

type route struct {
	id       string // added by the user
	method   string
	handler  func(http.ResponseWriter, *http.Request, *Router)
	//sections []*section
	url      string
}
