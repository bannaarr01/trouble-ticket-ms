package models

// Severity Indicate implication of the issue on the expected functionality e.g. of a system, application, service etc.
// e.g  Critical, Major, Minor.
type Severity struct {
	BaseModel
	Name string `gorm:"unique" json:"name"`
}
