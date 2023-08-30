package routes

import (
	"github.com/diiplexus/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBooksRoutes = func(router *mux.Router) {
	router.HandleFunc("/ping", controllers.Ping).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
