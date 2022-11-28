package service

import (
	"awesomeProject/internal/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var getPet = "https://petstore.swagger.io/v2/pet/findByStatus?status=available"

func GetPets() []models.Pet {
	resp, _ := http.Get(getPet)
	all, _ := ioutil.ReadAll(resp.Body)

	var petList []models.Pet

	json.Unmarshal(all, &petList)
	return petList
}
