package entities

import "time"

type Item struct {
	ID                int
	Title             string
	Description       string
	Url               string
	ImageFullUrl      string
	ImageThumbnailUrl string
	ExternalMetadata  map[string]interface{}
	CustomMetadata    map[string]interface{}
	Published         bool
	PublishedAt       time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
