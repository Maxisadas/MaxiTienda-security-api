package main

import (
	"api-rest/src/databases"
	"api-rest/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// routes
	routes.UserRoute(router);
	// Init DB
	databases.InitDB()
	// init server
	log.Fatal(http.ListenAndServe(":3000",router))
}
