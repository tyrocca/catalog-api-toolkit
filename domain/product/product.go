package product

import (
	entities "github.com/tyrocca/catalog-api-toolkit/entities"
)

type Product struct {
	item     *entities.Item
	variants *entities.Variant
}
