package models

import (
	"time"
	"trouble-ticket-ms/src/utils"
)

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

type BaseTroubleTicketDTO struct {
	Name                    string               `json:"name"`
	Description             string               `json:"description"`
	RequestedResolutionDate *time.Time           `json:"requested_resolution_date"`
	ExternalIdentifiers     []ExternalIdentifier `json:"external_identifiers"`
	RelatedEntities         []RelatedEntity      `json:"related_entities"`
	RelatedParties          []RelatedParty       `json:"related_parties"`
	StatusChanges           []StatusChange       `json:"status_changes"`
	Attachments             []Attachment         `json:"attachments"`
	Notes                   []Note               `json:"notes"`
}

type PartialTroubleTicketDTO struct {
	Ref         string `json:"ref"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TroubleTicketDTO struct {
	BaseModel
	PartialTroubleTicketDTO
	Type                   TypeDTO                 `json:"type"`
	Status                 StatusDTO               `json:"status"`
	Channel                ChannelDTO              `json:"channel"`
	Severity               SeverityDTO             `json:"severity"`
	Priority               PriorityDTO             `json:"priority"`
	ExternalIdentifiers    []ExternalIdentifierDTO `json:"external_identifiers"`
	RelatedEntities        []RelatedEntityDTO      `json:"related_entities"`
	RelatedParties         []RelatedPartyDTO       `json:"related_parties"`
	StatusChanges          []StatusChangeDTO       `json:"status_changes"`
	Attachments            []AttachmentDTO         `json:"attachments"`
	Notes                  []NoteDTO               `json:"notes"`
	ExpectedResolutionDate *time.Time              `json:"expected_resolution_date"`
	ResolutionDate         *time.Time              `json:"resolution_date"`
}

type CreateTroubleTicketDTO struct {
	Name                    string     `json:"name"`
	Description             string     `json:"description"`
	TypeID                  uint64     `json:"type_id"`
	ChannelID               uint64     `json:"channel_id"`
	RequestedResolutionDate *time.Time `json:"requested_resolution_date"`
}

type UpdateTroubleTicketDTO struct {
	CreateTroubleTicketDTO
	ExpectedResolutionDate time.Time `json:"expected_resolution_date"`
	ResolutionDate         time.Time `json:"resolution_date"`
}

func NewPartialTroubleTicketDTO(ref, name, description string) PartialTroubleTicketDTO {
	return PartialTroubleTicketDTO{ref, name, description}
}

// NewTroubleTicketDTO converts a TroubleTicket model to TroubleTicketDTO.
func NewTroubleTicketDTO(ticket *TroubleTicket) TroubleTicketDTO {
	return TroubleTicketDTO{
		PartialTroubleTicketDTO: NewPartialTroubleTicketDTO(ticket.Ref, ticket.Name, ticket.Description),
		ExpectedResolutionDate:  ticket.ExpectedResolutionDate,
		ResolutionDate:          ticket.ResolutionDate,

		BaseModel: NewBaseModel(ticket.BaseModel),
		Type:      NewTypeDTO(ticket.Type),
		Channel:   NewChannelDTO(ticket.Channel),
		Status:    NewStatusDTO(ticket.Status),
		Severity:  NewSeverityDTO(ticket.Severity),
		Priority:  NewPriorityDTO(ticket.Priority),

		ExternalIdentifiers: utils.TransformToDTO(ticket.ExternalIdentifiers, NewExternalIdentifierDTO),
		RelatedEntities:     utils.TransformToDTO(ticket.RelatedEntities, NewRelatedEntityDTO),
		RelatedParties:      utils.TransformToDTO(ticket.RelatedParties, NewRelatedPartyDTO),
		StatusChanges:       utils.TransformToDTO(ticket.StatusChanges, NewStatusChangeDTO),
		Attachments:         utils.TransformToDTO(ticket.Attachments, NewAttachmentDTO),
		Notes:               utils.TransformToDTO(ticket.Notes, NewNoteDTO),
	}
}

func NewTroubleTicket(
	c CreateTroubleTicketDTO,
	ref string,
	statusId, priorityId, severityId uint64,
	requestedResolutionDate *time.Time,
	expectedResolutionDate *time.Time,
	opts ...BaseModelOption,
) TroubleTicket {
	troubleTicket := TroubleTicket{
		Ref:                     ref,
		Name:                    c.Name,
		Description:             c.Description,
		TypeID:                  c.TypeID,
		StatusID:                statusId,
		ChannelID:               c.ChannelID,
		PriorityID:              priorityId,
		SeverityID:              severityId,
		RequestedResolutionDate: requestedResolutionDate,
		ExpectedResolutionDate:  expectedResolutionDate,
	}
	ApplyBaseMOptions(&troubleTicket.BaseModel, opts...)
	return troubleTicket
}
