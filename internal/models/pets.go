package models

type Pet struct {
	Id        int64         `json:"id"`
	Name      string        `json:"name"`
	PhotoUrls []interface{} `json:"photoUrls"`
	Tags      []interface{} `json:"tags"`
	Status    string        `json:"status"`
}
