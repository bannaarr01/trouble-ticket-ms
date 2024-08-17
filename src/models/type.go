package models

// Type Represent a business type of the trouble ticket e.g., incident, complain, request.
type Type struct {
	BaseModel
	Name string `gorm:"type:varchar(50);unique" json:"name"`
}

type TypeDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewTypeDTO(t Type) TypeDTO {
	return TypeDTO{
		ID:   t.ID,
		Name: t.Name,
	}
}
