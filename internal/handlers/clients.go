package handlers

import (
	"net/http"
	"testDataGoNew/internal/service"
)

func getClients(w http.ResponseWriter, r *http.Request) {
	service.GetClients(w, r)
}
