package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	SKU         string
	Price       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
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
