package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPMw ...
type HTTPMw func(http.HandlerFunc) http.HandlerFunc

// Router ...
type router interface {
	GET(pattern string, handler http.HandlerFunc, mw ...HTTPMw)
	POST(pattern string, handler http.HandlerFunc, mw ...HTTPMw)
	DELETE(pattern string, handler http.HandlerFunc, mw ...HTTPMw)
	PUT(pattern string, handler http.HandlerFunc, mw ...HTTPMw)

	GetServeHTTP() http.HandlerFunc
}

// HTTPRouter ...
type HTTPRouter struct {
	router *mux.Router
}

// NewRouter ...
// TODO: inject mux.Router methods in HTTPRouter
func NewRouter() *HTTPRouter {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()

	return &HTTPRouter{
		router: router,
	}
}
