package main

import (
	"github.com/gorilla/mux"
	rest2 "github.com/movie-search/Gateway/rest"
	"log"
	"net/http"
)

func main() {
	routes := rest2.InitPackage()

	//init mux routes
	router := mux.NewRouter()
	routes.InitRoutes(router)

	log.Printf("API Gateway up and running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
