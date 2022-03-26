package models

import (
	"log"
	"time"

	"github.com/tyrocca/catalog-api-toolkit/config"
)

type CatalogItem struct {
	Id                int       `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Url               string    `json:"url"`
	ImageFullUrl      string    `json:"image_full_url"`
	ImageThumbnailUrl string    `json:"image_thumbnail_url"`
	Published         bool      `json:"published"`
	PublishedAt       time.Time `json:"published_at"`
	CreatedAt         time.Time `json:"created"`
	UpdatedAt         time.Time `json:"updated"`
}
type CatalogItemCreateParams struct {
	Title             string
	Description       string
	Url               string
	ImageFullUrl      string
	ImageThumbnailUrl string
	PublishedAt       time.Time `json:"published_at"`
}

func (catalogItem *CatalogItem) Create(data CatalogItemCreateParams) (*CatalogItem, error) {
	var now = time.Now().UTC()
	var created_at = now
	var updated_at = now
	var published_at = now
	statement, _ := config.DB.Prepare("INSERT INTO catalog_item (title, description, url, image_full_url, image_thumbnail_url, published_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	result, err := statement.Exec(
		data.Title,
		data.Description,
		data.Url,
		data.ImageFullUrl,
		data.ImageThumbnailUrl,
		published_at,
		created_at,
		updated_at,
	)
	if err == nil {
		id, _ := result.LastInsertId()
		catalogItem.Id = int(id)
		catalogItem.Title = data.Title
		catalogItem.Description = data.Description
		catalogItem.Url = data.Url
		catalogItem.ImageFullUrl = data.ImageFullUrl
		catalogItem.PublishedAt = published_at
		catalogItem.Published = published_at.IsZero()
		catalogItem.CreatedAt = created_at
		catalogItem.UpdatedAt = updated_at

		return catalogItem, err
	}
	log.Println("Unable to create note", err.Error())
	return catalogItem, err
}
func (note *CatalogItem) GetAll() ([]CatalogItem, error) {
	rows, err := config.DB.Query("SELECT * FROM catalog_item")
	allCatalogItems := []CatalogItem{}
	if err == nil {
		for rows.Next() {
			var currentCatalogItem CatalogItem
			rows.Scan(
				&currentCatalogItem.Id,
				&currentCatalogItem.Title,
				&currentCatalogItem.Description,
				&currentCatalogItem.Url,
				&currentCatalogItem.ImageFullUrl,
				&currentCatalogItem.PublishedAt,
				&currentCatalogItem.CreatedAt,
				&currentCatalogItem.UpdatedAt,
			)
			currentCatalogItem.Published = currentCatalogItem.PublishedAt.IsZero()
			allCatalogItems = append(allCatalogItems, currentCatalogItem)
		}
		return allCatalogItems, err
	}
	return allCatalogItems, err
}

// func (note *Note) Fetch(id string) (*Note, error) {
// 	err := config.DB.QueryRow(
// 		"SELECT id, title, body, created_at, updated_at FROM notes WHERE id=?", id).Scan(
// 		&note.Id, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
// 	return note, err
// }
