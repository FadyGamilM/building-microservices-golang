package repository

import (
	"fmt"

	"github.com/FadyGamilM/product_api/data"
)

func GetAll() data.Products {
	return data.ProductList
}

func AddProduct(product *data.Product) {
	product.ID = data.ProductList[len(data.ProductList)-1].ID + 1
	data.ProductList = append(data.ProductList, product)
}

func UpdateProduct(productID int, newProduct *data.Product) error {
	for idx, prod := range data.ProductList {
		if prod.ID == productID {
			// data.ProductList = append(data.ProductList[:idx], data.ProductList[idx+1:]...)
			newProduct.ID = productID
			data.ProductList[idx] = newProduct
			return nil
		}
	}
	return fmt.Errorf("error while updating the products")
}
