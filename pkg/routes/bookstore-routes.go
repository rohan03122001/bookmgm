package routes

import (
	"github.com/gorilla/mux"
	"github.com/rohan03122001/bookmgm/pkg/controllers"
)

// RegisterBookStoreRoutes configures all API endpoints for the book store
var RegisterBookStoreRoutes = func(router *mux.Router) {
    // Define routes and their corresponding handlers
    router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
    router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
    router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
