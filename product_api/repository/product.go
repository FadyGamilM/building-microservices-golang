package repository

import (
	"github.com/FadyGamilM/product_api/data"
)

func GetAll() data.Products {
	return data.ProductList
}
