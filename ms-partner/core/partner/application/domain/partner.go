package domain

import "github.com/google/uuid"

type Partner struct {
	Identifier     uuid.UUID `json:"identifier"`
	Name           string    `json:"name"`
	PartnerType    string    `json:"partner_type"`
	UserIdentifier string    `json:"user_identifier"`
	Address        Address   `json:"address"`
	Menu           Menu      `json:"menu"`
}
