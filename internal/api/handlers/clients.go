package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testDataGoNew/internal/api/service"
	"testDataGoNew/internal/models"
)

func getClients(w http.ResponseWriter, r *http.Request) {
	service.GetClients(w, r)
}

func getClient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["idClient"])
	service.GetClientById(w, r, id)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["idClient"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	service.DeleteClientById(w, r, id)
}

func patch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["idClient"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	service.PatchClient(w, r, id)
}

func post(w http.ResponseWriter, r *http.Request) {
	var c models.Client
	body := r.Body

	all, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalf("Error - %v/\n", err.Error())
	}

	fmt.Println(string(all))

	err = json.Unmarshal(all, &c)
	if err != nil {
		log.Fatalf("Error - %v/\n", err.Error())
		w.WriteHeader(http.StatusForbidden)
	}

	service.PostClient(w, r, c)
}
