package entities

import "time"

type Item struct {
	ID                int
	Title             string
	Description       string
	Url               string
	ImageFullUrl      string
	ImageThumbnailUrl string
	Published         bool
	PublishedAt       time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
