package models

// Status A TroubleTicketStatusType. Possible values for the status of the trouble ticket e.g., pending, resolved etc.
type Status struct {
	BaseModel
	Name     string `json:"name"`
	Sequence int8   `json:"sequence"`
	Filter   int8   `json:"filter"`
}
