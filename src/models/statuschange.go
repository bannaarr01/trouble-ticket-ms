package models

import "time"

// StatusChange The status change history that are associated to the ticket
type StatusChange struct {
	BaseModel
	Reason          string `json:"reason"`
	StatusID        uint64 `gorm:"index;not null" json:"status_id"`
	Status          Status `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TroubleTicketID uint64 `gorm:"index;not null" json:"trouble_ticket_id"` // Fk for TroubleTicket
}

type StatusChangeDTO struct {
	ID        uint64     `json:"id"`
	Reason    string     `json:"reason"`
	Status    StatusDTO  `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

func NewStatusChangeDTO(statusChange StatusChange) StatusChangeDTO {
	return StatusChangeDTO{
		ID:        statusChange.ID,
		Reason:    statusChange.Reason,
		CreatedAt: statusChange.CreatedAt,
		CreatedBy: statusChange.CreatedBy,
		UpdatedAt: statusChange.UpdatedAt,
		UpdatedBy: statusChange.UpdatedBy,
		Status:    NewStatusDTO(statusChange.Status),
	}
}
