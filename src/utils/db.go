package utils

import "gorm.io/gorm"

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
