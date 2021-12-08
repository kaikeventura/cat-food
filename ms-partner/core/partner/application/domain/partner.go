package domain

type Partner struct {
	Identifier     string  `json:"identifier"`
	Name           string  `json:"name"`
	PartnerType    string  `json:"partner_type"`
	UserIdentifier string  `json:"user_identifier"`
	Address        Address `json:"address"`
}
