package models

type Client struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	ThirdName   string `json:"thirdName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	DateBirth   string `json:"dateBirth"`
	Reserve     bool   `json:"reserve"`
	Id          int    `json:"id"`
}
