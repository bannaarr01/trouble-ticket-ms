package models

type Party struct {
	BaseModel
	Name   string `gorm:"unique" json:"name"`
	Email  string `gorm:"unique" json:"email"`
	RoleID uint64 `gorm:"index;not null" json:"role_id"`
	Role   Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
