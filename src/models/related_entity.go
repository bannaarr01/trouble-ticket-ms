package models

// RelatedEntity in some way associated with the ticket, such as a bill, a product
type RelatedEntity struct {
	BaseModel
	Ref             string  `json:"ref"`
	Type            *string `json:"type" gorm:"type:varchar(20)"`
	Name            string  `json:"name" gorm:"type:varchar(50);not null"`
	Description     *string `json:"description" gorm:"type:text"`
	TroubleTicketID uint64  `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}

type RelatedEntityDTO struct {
	ID          uint64  `json:"id"`
	Ref         string  `json:"ref"`
	Type        *string `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func NewRelatedEntityDTO(relEntity RelatedEntity) RelatedEntityDTO {
	return RelatedEntityDTO{
		ID:          relEntity.ID,
		Ref:         relEntity.Ref,
		Type:        relEntity.Type,
		Name:        relEntity.Name,
		Description: relEntity.Description,
	}
}
