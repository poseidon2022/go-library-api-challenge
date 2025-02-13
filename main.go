package main

import (
	"log"
	"net/http"
	"go-library-api/src/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":3000", r))
}
