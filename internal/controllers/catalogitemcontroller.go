package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tyrocca/catalog-api-toolkit/internal/models"
)

type CatalogItemController struct{}

func (_ *CatalogItemController) CreateNew(c echo.Context) {
	var params models.CatalogItemCreateParams
	var catalogItem models.CatalogItem
	err := c.Bind(params)
	if err == nil {
		_, creationError := catalogItem.Create(params)
		if creationError == nil {
			c.JSON(http.StatusCreated, {
				"message": "CatalogItem created successfully",
				"note":    CatalogItem,
			})
		} else {
			c.String(http.StatusInternalServerError, creationError.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// func (_ *CatalogItemController) GetAllNotes(c *gin.Context) {
// 	var note models.Note
// 	notes, err := note.GetAll()
// 	if err == nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "All Notes",
// 			"notes":   notes,
// 		})
// 	} else {
// 		c.String(http.StatusInternalServerError, err.Error())
// 	}
// }
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
