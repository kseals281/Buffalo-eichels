package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sabrestats/middleware"
)

// TODO: Update all stats once per day
// TODO: Setup database to house all player information
// TODO: Rebuild connection to React with the correct models

func main() {
	r := mux.NewRouter()

	//r.HandleFunc("/api/update")

	r.HandleFunc("/api/forwards", middleware.Forwards).Methods("GET")
	r.HandleFunc("/api/defencemen", middleware.Defencemen).Methods("GET")
	//r.HandleFunc("/api/goalies", middleware.Goalies).Methods("GET")

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
