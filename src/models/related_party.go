package models

type RelatedParty struct {
	BaseModel
	PartyID         uint64 `gorm:"index;not null" json:"role_id"`
	Party           Party  `gorm:"foreignKey:PartyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}
