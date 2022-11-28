package handlers

import (
	"awesomeProject/internal/service"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	service.GetHome(w, r)
}
