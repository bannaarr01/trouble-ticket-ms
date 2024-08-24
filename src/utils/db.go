package utils

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// NestedPreload recursively applies GORM preloads for nested relationships.
// Accepts a variadic number of fields to preload in sequence.
//
// Returns: func(db *gorm.DB) *gorm.DB: A function that applies nested preloads.
func NestedPreload(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) == 0 {
			return db
		}

		// Preload the first field and recursively preload the rest
		return db.Preload(fields[0], NestedPreload(fields[1:]...))
	}
}

// CheckRelatedRecordExists checks if a related record exists in the database.
//
// It checks if a record exists in the database based on the provided model, ID, and field name.
// It returns an error if the record does not exist, or nil if it does.
func CheckRelatedRecordExists(tx *gorm.DB, model interface{}, id uint64, fieldName string) error {
	var count int64
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	modelName := strings.ToLower(t.Name())

	if err := tx.Model(model).Where(fieldName+" = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("%s with ID %d does not exist", modelName, id)
	}

	return nil
}
