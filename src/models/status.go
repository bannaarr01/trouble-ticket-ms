package models

// Status A TroubleTicketStatusType. Possible values for the status of the trouble ticket e.g., pending, resolved etc.
type Status struct {
	BaseModel
	Name     string `json:"name"`
	Sequence int8   `json:"sequence"`
	Filter   int8   `json:"filter"`
}

type StatusDTO struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Sequence int8   `json:"sequence"`
	Filter   int8   `json:"filter"`
}

func NewStatusDTO(status Status) StatusDTO {
	return StatusDTO{
		ID:       status.ID,
		Name:     status.Name,
		Sequence: status.Sequence,
		Filter:   status.Filter,
	}
}

func NewStatus(name string, sequence, filter int8, opts ...BaseModelOption) Status {
	status := Status{
		BaseModel: BaseModel{},
		Name:      name,
		Sequence:  sequence,
		Filter:    filter,
	}
	ApplyBaseMOptions(&status.BaseModel, opts...)
	return status
}
