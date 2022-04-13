package company_catalog

import (
	product "github.com/tyrocca/catalog-api-toolkit/domain/product"
	entities "github.com/tyrocca/catalog-api-toolkit/entities"
)

type CompanyCatalog struct {
	catalog    *entities.Catalog
	company    *entities.CompanyCatalog
	categories *entities.Category
	products   *product.Product
}
