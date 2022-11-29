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
			Id:         1,
			Username:   "Maks",
			FirstName:  "Hor",
			LastName:   "Hor",
			Email:      "m@m.ru",
			Password:   "123",
			Phone:      "+79991112233",
			UserStatus: 0,
		},
		{
			Id:         2,
			Username:   "Ivan",
			FirstName:  "Petr",
			LastName:   "Kold",
			Email:      "m@dd.ru",
			Password:   "123",
			Phone:      "+79991112233",
			UserStatus: 0,
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
