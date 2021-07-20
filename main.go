package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi.go/handlers"
	

	"github.com/gorilla/mux"
)
	
type home struct {
	Message string `json:"message"`
}

func homelink(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(home{Message: "Welcome"})
	
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homelink).Methods("GET")
	router.HandleFunc("/products", handlers.Getallproducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
	

}




