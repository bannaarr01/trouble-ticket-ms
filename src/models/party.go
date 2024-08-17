package models

type Party struct {
	BaseModel
	Name   string `gorm:"unique" json:"name"`
	Email  string `gorm:"unique" json:"email"`
	RoleID uint64 `gorm:"index;not null" json:"role_id"`
	Role   Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PartyDTO struct {
	ID    uint64  `json:"id"`
	Name  string  `son:"name"`
	Email string  `json:"email"`
	Role  RoleDTO `json:"role"`
}

func NewPartyDTO(party Party) PartyDTO {
	return PartyDTO{
		ID:    party.ID,
		Name:  party.Name,
		Email: party.Email,
		Role:  NewRoleDTO(party.Role),
	}
}
