package models

import "trouble-ticket-ms/src/utils"

type Filters struct {
	Types      []Type
	Statuses   []Status
	Severities []Severity
	Channels   []Channel
	Priorities []Priority
	Roles      []Role
}

type FiltersDTO struct {
	Types      []TypeDTO     `json:"types"`
	Statuses   []StatusDTO   `json:"statuses"`
	Severities []SeverityDTO `json:"severities"`
	Channels   []ChannelDTO  `json:"channels"`
	Priorities []PriorityDTO `json:"priorities"`
	Roles      []RoleDTO     `json:"roles"`
}

func NewFilterDTO(filters Filters) FiltersDTO {
	return FiltersDTO{
		Types:      utils.TransformToDTO(filters.Types, NewTypeDTO),
		Statuses:   utils.TransformToDTO(filters.Statuses, NewStatusDTO),
		Severities: utils.TransformToDTO(filters.Severities, NewSeverityDTO),
		Channels:   utils.TransformToDTO(filters.Channels, NewChannelDTO),
		Priorities: utils.TransformToDTO(filters.Priorities, NewPriorityDTO),
		Roles:      utils.TransformToDTO(filters.Roles, NewRoleDTO),
	}
}
