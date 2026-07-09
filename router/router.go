package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// route struct containing data about route, like middleware, path etc..
type Route struct {
	Path       string
	Method     string
	Middleware []func(http.ResponseWriter, *http.Request)
	Handler    func(http.ResponseWriter, *http.Request)
}

// route struct factory to set route parameters
func NewRoute() *Route {
	return &Route{}
}

// middleware method for route
func (route *Route) HandleMiddleWare(w http.ResponseWriter, r *http.Request) {

	log.Printf("%v %v", route.Method, route.Path)
	//check for middleware and execute
	for _, middleware := range route.Middleware {
		middleware(w, r)
	}

	//handle request
	route.Handler(w, r)

}

// router struct to hold all routes
type Router struct {
	Route map[string]*Route
}

// router factory
func NewRouter() *Router {
	return &Router{
		Route: make(map[string]*Route),
	}
}

// get route setter
func (Router *Router) GET(path string, handler func(http.ResponseWriter, *http.Request)) *Route {
	//create route instance
	Route := NewRoute()
	Route.Path = path
	Route.Method = "GET"
	// Route.Middleware = []func(http.ResponseWriter, *http.Request){}
	Route.Handler = handler

	//create key for map lookup
	key := "GET" + Route.Path
	fmt.Println(key)
	Router.Route[key] = Route

	//return route pointer
	return Route
}

// implements handler interface
func (Router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check request method

	switch r.Method {
	case "GET":
		if route, exists := Router.Route["GET"+strings.TrimSuffix(r.URL.Path, "/")]; exists {
			//run potential middleware
			route.HandleMiddleWare(w, r)
			// route.Handler(w, r)
		} else {
			w.WriteHeader(404)
			fmt.Fprint(w, "404 not found")
		}
	default:

	}
}
