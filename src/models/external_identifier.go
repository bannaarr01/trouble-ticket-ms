package models

// ExternalIdentifier If originated from a different system such as a product order, handed off from a commerce platform
// into an order handling system multiple external IDs can be held for a single entity, e.g. if the entity passed
// through multiple systems on the way to the current system. In this case the consumer is expected to sequence the IDs
// in the array in reverse  order of provenance
type ExternalIdentifier struct {
	BaseModel
	Owner           string `json:"owner"`
	Ref             string `json:"ref"` // e.g order_id n so on
	TypeID          uint64 `gorm:"index;not null" json:"ticket_type_id"`
	Type            Type   `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}

type ExternalIdentifierDTO struct {
	ID    uint64  `json:"id"`
	Owner string  `json:"owner"`
	Ref   string  `json:"ref"`
	Type  TypeDTO `json:"type"`
}

func NewExternalIdentifierDTO(extId ExternalIdentifier) ExternalIdentifierDTO {
	return ExternalIdentifierDTO{
		ID:    extId.ID,
		Owner: extId.Owner,
		Ref:   extId.Ref,
		Type:  NewTypeDTO(extId.Type),
	}
}
