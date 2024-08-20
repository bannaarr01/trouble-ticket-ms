package models

// Channel The channel to which the resource reference to. e.g., Sales, Support, Billing, IT, HR, Finance, Operations...
type Channel struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null" json:"name"`
}

type ChannelDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewChanel(channel Channel) ChannelDTO {
	return ChannelDTO{
		ID:   channel.ID,
		Name: channel.Name,
	}
}

func NewChannel(name string, opts ...BaseModelOption) Channel {
	channel := Channel{BaseModel: BaseModel{}, Name: name}
	ApplyBaseMOptions(&channel.BaseModel, opts...)
	return channel
}
