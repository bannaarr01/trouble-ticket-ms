package models

import (
	"reflect"
	"time"
)

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

// SetField sets the value of a field in a BaseModel.
// Args:
//
//	fieldName: The name of the field to set.
//	value: The value to set the field to.
//
// Returns:
//
//	A BaseModelOption that sets the field when applied to a BaseModel.
func SetField(fieldName string, value interface{}) BaseModelOption {
	return func(bm *BaseModel) {
		v := reflect.ValueOf(bm).Elem()
		if v.IsValid() {
			field := v.FieldByName(fieldName)
			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(value))
			}
		}
	}
}
