package worldnews

import "time"

type WorldNews struct {
	ImageURL      *string    `json:"imageUrl,omitempty"`
	Title         *string    `json:"title,omitempty"`
	GMTTime       *time.Time `json:"gmtTime,omitempty"`
	SourceStr     *SourceStr `json:"sourceStr,omitempty"`
	Lead          *string    `json:"lead,omitempty"`
	SourceIconURL *string    `json:"sourceIconUrl,omitempty"`
	Page          *Page      `json:"page,omitempty"`
}

type Page struct {
	URL *string `json:"url,omitempty"`
}

type SourceStr string

const (
	FotMob SourceStr = "FotMob"
)
