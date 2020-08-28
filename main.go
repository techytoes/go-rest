package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
		"message": "🚀🌟🌈",
	})
}

func main() {

	// init router
	r := mux.NewRouter()

	// Routes for the various methods and resources.
	r.HandleFunc("/", Home).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
