package models

type RelatedParty struct {
	BaseModel
	PartyID         uint64 `gorm:"index;not null" json:"role_id"`
	Party           Party  `gorm:"foreignKey:PartyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}

type RelatedPartyDTO struct {
	ID    uint64   `json:"id"`
	Party PartyDTO `json:"party"`
}

func NewRelatedPartyDTO(relParty RelatedParty) RelatedPartyDTO {
	return RelatedPartyDTO{
		ID:    relParty.ID,
		Party: NewPartyDTO(relParty.Party),
	}
}
