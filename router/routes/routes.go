package routes

import (
	"astro/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct
type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Load is  a func
func Load() []Route {
	routes := AstroRoutes
	routes = append(routes, BannerOfferRoutes...)

	return routes
}

//SetUpRoutes is a func
func SetUpRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}

//SetUpRoutesWithMiddlewares is  a func
func SetUpRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(route.Handler))).Methods(route.Method)
	}
	return r
}
