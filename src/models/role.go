package models

// Role Such as initiator, customer, salesAgent, user.
type Role struct {
	BaseModel
	Name     string `json:"name"`
	Sequence int8   `json:"sequence"`
	Filter   int8   `json:"filter"`
}

type RoleDTO struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Sequence int8   `json:"sequence"`
	Filter   int8   `json:"filter"`
}

func NewRoleDTO(role Role) RoleDTO {
	return RoleDTO{
		ID:       role.ID,
		Name:     role.Name,
		Sequence: role.Sequence,
		Filter:   role.Filter,
	}
}
