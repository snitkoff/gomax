package gomax

import "net/http"

type Router struct {
	routes []*route
	index  *index
	url    string
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if route := r.index.find(req); route != nil {
		route.handler(w, req, r)
	} else {
		r.Default(w, req)
	}
}

func (r *Router) Default(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
}
