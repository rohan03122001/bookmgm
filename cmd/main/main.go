package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohan03122001/bookmgm/pkg/routes"
)

// main initializes and starts the HTTP server with configured routes
func main() {
    // Initialize router with Gorilla Mux
    r := mux.NewRouter()
    
    // Register API routes from the routes package
    routes.RegisterBookStoreRoutes(r)
    
    // Set up default route handler
    http.Handle("/", r)
    
    // Start server on port 9010 and log any fatal errors
    log.Fatal(http.ListenAndServe("localhost:9010", r))
}