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
	Ref                     string               `json:"ref"`
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
	BaseTroubleTicketDTO
	TypeID     uint64 `json:"type_id"`
	StatusID   uint64 `json:"status_id"`
	ChannelID  uint64 `json:"channel_id"`
	PriorityID uint64 `json:"priority_id"`
	SeverityID uint64 `json:"severity_id"`
}

type UpdateTroubleTicketDTO struct {
	CreateTroubleTicketDTO
	ExpectedResolutionDate time.Time `json:"expected_resolution_date"`
	ResolutionDate         time.Time `json:"resolution_date"`
}

type TroubleTicketsDTO struct {
	TroubleTickets []TroubleTicketDTO `json:"trouble_tickets"`
}

func NewPartialTroubleTicketDTO(ref, name, description string) PartialTroubleTicketDTO {
	return PartialTroubleTicketDTO{ref, name, description}
}

// NewTroubleTicketDTO converts a TroubleTicket model to TroubleTicketDTO.
func NewTroubleTicketDTO(ticket TroubleTicket) TroubleTicketDTO {
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

		ExternalIdentifiers: transformToExternalIdentifierDTO(ticket.ExternalIdentifiers),
		RelatedEntities:     transformToRelatedEntityDTO(ticket.RelatedEntities),
		RelatedParties:      transformToRelatedPartyDTO(ticket.RelatedParties),
		StatusChanges:       transformToStatusChangeDTO(ticket.StatusChanges),
		Attachments:         transformToAttachmentDTO(ticket.Attachments),
		Notes:               transformToNoteDTO(ticket.Notes),
	}
}

// transformToNoteDTO converts Note model to []NoteDTO.
func transformToNoteDTO(notes []Note) []NoteDTO {
	return utils.SerializeModelToDTO(notes, func(note Note) NoteDTO {
		return NewNoteDTO(note)
	})
}

// transformToAttachmentDTO converts Attachment model to []AttachmentDTO.
func transformToAttachmentDTO(attachments []Attachment) []AttachmentDTO {
	return utils.SerializeModelToDTO(attachments, func(attachment Attachment) AttachmentDTO {
		return NewAttachmentDTO(attachment)
	})
}

// transformToStatusChangeDTO converts StatusChange model to []StatusChangeDTO.
func transformToStatusChangeDTO(statusChanges []StatusChange) []StatusChangeDTO {
	return utils.SerializeModelToDTO(statusChanges, func(statusChange StatusChange) StatusChangeDTO {
		return NewStatusChangeDTO(statusChange)
	})
}

// transformToExternalIdentifierDTO converts RelatedEntity model to []RelatedEntityDTO.
func transformToRelatedEntityDTO(relEntities []RelatedEntity) []RelatedEntityDTO {
	return utils.SerializeModelToDTO(relEntities, func(relEntity RelatedEntity) RelatedEntityDTO {
		return NewRelatedEntityDTO(relEntity)
	})
}

// transformToExternalIdentifierDTO converts RelatedParty model to []RelatedPartyDTO.
func transformToRelatedPartyDTO(relParties []RelatedParty) []RelatedPartyDTO {
	return utils.SerializeModelToDTO(relParties, func(relParty RelatedParty) RelatedPartyDTO {
		return NewRelatedPartyDTO(relParty)
	})
}

// transformToExternalIdentifierDTO converts ExternalIdentifier model to []ExternalIdentifierDTO.
func transformToExternalIdentifierDTO(extIdentifiers []ExternalIdentifier) []ExternalIdentifierDTO {
	return utils.SerializeModelToDTO(extIdentifiers, func(extId ExternalIdentifier) ExternalIdentifierDTO {
		return NewExternalIdentifierDTO(extId)
	})
}
