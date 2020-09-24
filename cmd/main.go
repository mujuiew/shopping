package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mujuiew/api-shopping/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", api.Reister).Methods("POST")
	r.HandleFunc("/account/store", api.Store)

	http.ListenAndServe(":8080", r)
}
