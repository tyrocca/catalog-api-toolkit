package controllers

import (
	"database/sql"
)

type BaseController struct {
	db *sql.DB
}

type ControllerRegistery struct {
	CatalogItemController *CatalogItemController
}

func CreateControllers(db *sql.DB) *ControllerRegistery {
	return &ControllerRegistery{
		CatalogItemController: GetCatalogItemController(db),
	}
}
