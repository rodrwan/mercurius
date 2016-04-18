package server

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Method",
		"POST",
		"/scraping/{type}",
		Method,
	},
	Route{
		"PayloadHandler",
		"POST",
		"/payload",
		PayloadHandler,
	},
}
