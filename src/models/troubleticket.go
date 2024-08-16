package models

import "time"

type TroubleTicket struct {
	BaseModel
	Ref                     string               `gorm:"type:varchar(20);unique;not null" json:"ref"`
	Name                    string               `gorm:"type:varchar(20);not null" json:"name"`
	Description             string               `gorm:"type:text;not null" json:"description"`
	RequestedResolutionDate *time.Time           `json:"requested_resolution_date"` // request by customer
	ExpectedResolutionDate  *time.Time           `json:"expected_resolution_date"`  // determined by the sys. So default null
	ResolutionDate          *time.Time           `json:"resolution_date"`           // DateTime the ticket was resolved
	TypeID                  uint64               `gorm:"index;not null" json:"type_id"`
	StatusID                uint64               `gorm:"index;not null" json:"status_id"`
	ChannelID               uint64               `gorm:"index;not null" json:"channel_id"`
	PriorityID              uint64               `gorm:"index;not null" json:"priority_id"` // Fk for Priority
	SeverityID              uint64               `gorm:"index;not null" json:"severity_id"`
	ExternalIdentifiers     []ExternalIdentifier `gorm:"foreignKey:TroubleTicketID" json:"external_identifiers"`
	RelatedEntities         []RelatedEntity      `gorm:"foreignKey:TroubleTicketID" json:"related_entities"`
	RelatedParties          []RelatedParty       `gorm:"foreignKey:TroubleTicketID" json:"related_parties"`
	StatusChanges           []StatusChange       `gorm:"foreignKey:TroubleTicketID" json:"status_changes"`
	Attachments             []Attachment         `gorm:"foreignKey:TroubleTicketID" json:"attachments"`
	Notes                   []Note               `gorm:"foreignKey:TroubleTicketID" json:"notes"`
	// Establish the relationship
	Type     Type     `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status   Status   `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Channel  Channel  `gorm:"foreignKey:ChannelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Severity Severity `gorm:"foreignKey:SeverityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Priority Priority `gorm:"foreignKey:PriorityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
