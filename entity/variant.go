package entities

import "time"

type Variant struct {
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
