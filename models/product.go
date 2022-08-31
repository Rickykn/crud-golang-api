package models

import (
	"github.com/Rickykn/crud-golang-api.git/database"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

func GetProducts() ([]*Product, error) {
	db := database.Get()

	rows, err := db.Query("SELECT id,name,description,quantity FROM products_tab")

	if err != nil {
		return nil, err
	}

	products := []*Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func GetProductById(id int) (*Product, error) {
	db := database.Get()

	product := &Product{}

	err := db.QueryRow("SELECT id,name,description,quantity FROM products_tab WHERE id = $1", id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return product, nil

}
