package models

import "time"

type BaseModel struct {
	ID        uint64     `gorm:"primaryKey" json:"id"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"` // autoPopulate with the current timestamp on record creation
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`
}
