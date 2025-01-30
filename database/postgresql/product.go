package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func CreateProduct(product config.Product) (int, error) {
	var id int
	query := `INSERT INTO products (name, description, price, quantity, seller_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := database.PsqlDB.QueryRow(query, product.Name, product.Description, product.Price, product.Quantity, product.SellerID).Scan(&id)
	return id, err
}

func GetProductByID(id int) (config.Product, error) {
	var product config.Product
	query := `SELECT * FROM products WHERE id = $1`
	err := database.PsqlDB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.SellerID)
	return product, err
}

func GetProducts() ([]config.Product, error) {
	var products []config.Product

	query := `SELECT * FROM products`
	rows, err := database.PsqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product config.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.SellerID)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProduct(product config.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, quantity = $4, seller_id = $5 WHERE id = $6`
	_, err := database.PsqlDB.Exec(query, product.Name, product.Description, product.Price, product.Quantity, product.SellerID, product.ID)
	return err
}

func DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := database.PsqlDB.Exec(query, id)
	return err
}
