package router

import (
	"log"
	"net/http"
	"time"
)

// ChainsMiddlewares ...
func (h *HTTPRouter) ChainsMiddlewares(method string, handler http.HandlerFunc, mw ...HTTPMw) http.HandlerFunc {
	mws := make([]HTTPMw, 0)
	mws = append(mws, mw...)

	for _, httpMw := range mws {
		handler = httpMw(handler)
	}

	return handler
}

// Handle ...
func (h *HTTPRouter) Handle(method string, pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	handler = h.ChainsMiddlewares(method, handler, mw...)
	h.router.HandleFunc(pattern, handler).Methods(method, "OPTIONS")
}

func (h *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "OPTIONS" {
		log.Printf("\033[0;32m%s\033[0m \033[0;36m%s\033[0m %s",
			r.Method,
			r.RequestURI,
			time.Since(time.Now()),
		)
	}

	h.router.ServeHTTP(w, r)
}

// GET ...
func (h *HTTPRouter) GET(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("GET", pattern, handler, mw...)
}

// POST ...
func (h *HTTPRouter) POST(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("POST", pattern, handler, mw...)
}

// PUT ...
func (h *HTTPRouter) PUT(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("PUT", pattern, handler, mw...)
}

// DELETE ...
func (h *HTTPRouter) DELETE(pattern string, handler http.HandlerFunc, mw ...HTTPMw) {
	h.Handle("DELETE", pattern, handler, mw...)
}
