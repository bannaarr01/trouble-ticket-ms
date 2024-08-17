package models

// Priority indicate how quickly the issue should be resolved. e.g., Critical, High, Medium, Low.
// Also considering the severity, ticket type etc
type Priority struct {
	BaseModel
	Type     string `json:"type"`
	Sequence int8   `json:"sequence"`
}

type PriorityDTO struct {
	ID       uint64 `json:"id"`
	Type     string `json:"type"`
	Sequence int8   `json:"sequence"`
}

func NewPriorityDTO(priority Priority) PriorityDTO {
	return PriorityDTO{
		ID:       priority.ID,
		Type:     priority.Type,
		Sequence: priority.Sequence,
	}
}
