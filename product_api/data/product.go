package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`
}

type Products []*Product

func GetProducts() []*Product {
	return ProductList
}

func AddProduct(product *Product) {
	product.ID = ProductList[len(ProductList)-1].ID + 1
	ProductList = append(ProductList, product)
}

// a method that extract json data format from the given data type [Products]
func (products *Products) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(products)
}

// a method extract the Product data format fromt he json format
func (product *Product) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&product)
}

// memic the database
var ProductList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "fratby coffe",
		SKU:         "latte_1_coffe",
		Price:       29.99,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Ice Mocha",
		Description: "mocha with ice cubes and milk with a shot of coffe",
		SKU:         "mocka_2_coffe",
		Price:       35.30,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
