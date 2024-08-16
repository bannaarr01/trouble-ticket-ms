package models

// Type Represent a business type of the trouble ticket e.g., incident, complain, request.
type Type struct {
	BaseModel
	Name string `gorm:"type:varchar(50);unique"`
}
