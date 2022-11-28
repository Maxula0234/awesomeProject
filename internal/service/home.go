package service

import (
	"encoding/json"
	"log"
	"net/http"
	"testDataGoNew/version"
	"time"
)

func GetHome(w http.ResponseWriter, _ *http.Request) {
	info := struct {
		BuildTime time.Time `json:"buildTime"`
		Commit    string    `json:"commit"`
	}{
		version.BuildTime, version.Commit,
	}

	body, err := json.Marshal(info)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
