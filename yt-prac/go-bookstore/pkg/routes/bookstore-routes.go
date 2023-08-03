package routes

import (
	"github.com/gorilla/mux"
	"github.com/raphhawk/bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST") //OK
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET") //OK
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")  //OK
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
