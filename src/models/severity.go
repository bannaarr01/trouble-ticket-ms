package models

// Severity Indicate implication of the issue on the expected functionality e.g. of a system, application, service etc.
// e.g  Critical, Major, Minor.
type Severity struct {
	BaseModel
	Name string `gorm:"unique" json:"name"`
}

type SeverityDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewSeverityDTO(severity Severity) SeverityDTO {
	return SeverityDTO{
		ID:   severity.ID,
		Name: severity.Name,
	}
}

func NewSeverity(name string, opts ...BaseModelOption) Severity {
	severity := Severity{BaseModel: BaseModel{}, Name: name}
	ApplyBaseMOptions(&severity.BaseModel, opts...)
	return severity
}
