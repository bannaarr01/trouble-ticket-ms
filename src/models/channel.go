package models

// Channel The channel to which the resource reference to. e.g., Sales, Support, Billing, IT, HR, Finance, Operations...
type Channel struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null" json:"name"`
}
