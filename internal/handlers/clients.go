package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"testDataGoNew/internal/service"
)

func getClients(w http.ResponseWriter, r *http.Request) {
	service.GetClients(w, r)
}

func getClient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	service.GetClientById(w, r, id)
}
