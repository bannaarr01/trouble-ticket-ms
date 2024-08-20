package models

import "time"

type BaseModelOption func(model *BaseModel)

type BaseModel struct {
	ID        uint64     `gorm:"primaryKey" json:"id"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"` // autoPopulate with the current timestamp on record creation
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`
}

func NewBaseModel(bm BaseModel) BaseModel {
	return BaseModel{
		ID:        bm.ID,
		CreatedBy: bm.CreatedBy,
		CreatedAt: bm.CreatedAt,
		UpdatedAt: bm.UpdatedAt,
		UpdatedBy: bm.UpdatedBy,
		DeletedAt: bm.DeletedAt,
		DeletedBy: bm.DeletedBy,
	}
}

// ApplyBaseMOptions to apply Base Model fields if needed
func ApplyBaseMOptions(target *BaseModel, opts ...BaseModelOption) {
	for _, opt := range opts {
		opt(target)
	}
}
