package domain

import "github.com/google/uuid"

type Menu struct {
	Identifier        uuid.UUID `json:"identifier"`
	PartnerIdentifier string    `json:"partner_identifier"`
}
