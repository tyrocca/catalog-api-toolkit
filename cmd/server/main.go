package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tyrocca/catalog-api-toolkit/config"
	"github.com/tyrocca/catalog-api-toolkit/internal/controllers"
)

func main() {
	db, initialized, err := config.InitializeDB()
	if err != nil {
		log.Fatalln("Sql Error", err.Error())
	} else if !initialized {
		log.Fatalln("Db not initialized")
	}
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// var catalogItemController

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	controller := controllers.CreateControllers(db)

	// Catalog item operations
	e.GET("/catalog-items", controller.CatalogItemController.GetAll)
	e.POST("/catalog-items", controller.CatalogItemController.CreateNew)

	e.Logger.Fatal(e.Start(":1323"))
}
