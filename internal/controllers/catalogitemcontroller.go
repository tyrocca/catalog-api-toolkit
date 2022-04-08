package controllers

import (
	"net/http"

	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/tyrocca/catalog-api-toolkit/internal/models"
)

type CatalogItemController struct {
	db *sql.DB
}

func GetCatalogItemController(db *sql.DB) *CatalogItemController {
	return &CatalogItemController{
		db: db,
	}
}

func (controller *CatalogItemController) CreateNew(c echo.Context) error {
	var params models.CatalogItemCreateParams
	var catalogItem models.CatalogItem
	err := c.Bind(params)
	if err == nil {
		_, creationError := catalogItem.Create(controller.db, params)
		if creationError == nil {
			c.JSON(http.StatusCreated, catalogItem)
		} else {
			c.String(http.StatusInternalServerError, creationError.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return err
}

func (controller *CatalogItemController) GetAll(c echo.Context) error {
	var item models.CatalogItem
	items, err := item.GetAll(controller.db)
	if err == nil {
		c.JSON(http.StatusOK, items)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return err
}

// func (_ *CatalogItemController) GetSingleNote(c *gin.Context) {
// 	var note models.Note
// 	id := c.Param("note_id")
// 	_, err := note.Fetch(id)
// 	if err == nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Single Note",
// 			"note":    note,
// 		})
// 	} else {
// 		c.String(http.StatusInternalServerError, err.Error())
// 	}
// }
