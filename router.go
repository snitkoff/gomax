package gomax

import "net/http"

// Router is a main router structure
type Router struct {
	routes []*route
	index  *index
	url    string
}

// NewRouter using for creating new empty rouret structure
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

// Default is a router method to return 404 response
func (r *Router) Default(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
}
