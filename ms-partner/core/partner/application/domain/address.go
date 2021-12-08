package domain

type Address struct {
	Identifier string `json:"identifier"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}
