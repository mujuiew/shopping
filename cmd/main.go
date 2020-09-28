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

	r.HandleFunc("/account/addproduct", api.AddProduct).Methods("POST")

	http.ListenAndServe(":8080", r)
}
