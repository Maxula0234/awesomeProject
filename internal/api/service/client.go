package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testDataGoNew/internal/models"
)

var (
	clients = []models.Client{
		{
			FirstName:   "Иван",
			LastName:    "Иванов",
			ThirdName:   "Иванович",
			PhoneNumber: "79991112233",
			Email:       "m@test.ru",
			DateBirth:   "string",
			Reserve:     false,
			Id:          0,
		},
		{
			FirstName:   "Максим",
			LastName:    "Хор",
			ThirdName:   "Иванович",
			PhoneNumber: "79991112233",
			Email:       "m@tвыst.ru",
			DateBirth:   "string",
			Reserve:     false,
			Id:          0,
		},
	}
)

func GetClint(id int) models.Client {
	fmt.Printf("Нашли клиента %s\n", id)
	return models.Client{}
}

func GetClients(w http.ResponseWriter, _ *http.Request) []models.Client {
	fmt.Printf("Нашли много клиентов\n")

	body, err := json.Marshal(clients)
	if err != nil {
		log.Fatalf("Error - %v/\n", err.Error())
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return []models.Client{}
}

func GetClientById(w http.ResponseWriter, _ *http.Request, id int) {
	var c models.Client
	for _, client := range clients {
		if client.Id == id {
			c = client
		}
	}

	if c.Id != id {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	body, err := json.Marshal(c)
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func DeleteClientById(w http.ResponseWriter, _ *http.Request, id int) {
	w.WriteHeader(http.StatusOK)
}

func PatchClient(w http.ResponseWriter, _ *http.Request, id int) models.Client {
	return models.Client{}
}

func PostClient(w http.ResponseWriter, _ *http.Request, c models.Client) {
	clients = append(clients, c)

	marshal, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Error - %v/\n", err.Error())
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(marshal)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusForbidden)
	}

	fmt.Printf("Клиентов стало - %v\n", len(clients))

}
