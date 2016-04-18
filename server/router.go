package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	bearer := fmt.Sprintf("Bearer %s", os.Getenv("BEARER"))

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		// "Accept", "application/vnd.finciero+json; version=1").
		router.
			Headers("Content-Type", "application/json",
			"Authorization", bearer).
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
