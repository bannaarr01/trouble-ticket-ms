package models

// Role Such as initiator, customer, salesAgent, user.
type Role struct {
	BaseModel
	Name     string `json:"name"`
	Sequence int8   `json:"sequence"`
	Filter   int8   `json:"filter"`
}
