package domain

type UserDomain struct {
	Identifier string    `json:"identifier"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	Document   string    `json:"document"`
	UserType   string    `json:"user_type"`
	Address    []Address `json:"addresses"`
}
