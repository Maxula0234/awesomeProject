package handlers

import (
	"net/http"
	"testDataGoNew/internal/api/service"
)

func home(w http.ResponseWriter, r *http.Request) {
	service.GetHome(w, r)
}
