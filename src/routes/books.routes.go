package router

import (
	"github.com/gorilla/mux"
	"go-library-api/src/controllers"
)

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
