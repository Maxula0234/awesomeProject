package handlers

import (
	"awesomeProject/internal/service"
	"net/http"
)

func getClients(w http.ResponseWriter, r *http.Request) {
	service.GetClients(w, r)
}
