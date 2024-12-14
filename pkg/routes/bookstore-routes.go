package routes

import (
	"github.com/gorilla/mux"
	"github.com/rohan03122001/bookmgm/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/",controllers.DeleteBook).Methods("DELETE")
}
