package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/clients", getClients).Methods("GET")
	r.HandleFunc("/clients/{id}", getClient).Methods("GET")
	r.HandleFunc("/clients", post).Methods("POST")
	return r

}
