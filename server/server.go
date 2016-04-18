package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// API ...
func API(port string) {
	router := NewRouter()
	fmt.Printf("Running on port: %s\n", port)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}
