package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func CreateProduct(product config.Product) (int, error) {
	var id int
	query := `INSERT INTO products (name, description, price, seller_id) VALUES ($1, $2, $3) RETURNING id`
	err := database.PsqlDB.QueryRow(query, product.Name, product.Description, product.Price, product.SellerID).Scan(&id)
	return id, err
}

func GetProductByID(id int) (config.Product, error) {
	var product config.Product
	query := `SELECT id, name, description, price, seller_id FROM products WHERE id = $1`
	err := database.PsqlDB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SellerID)
	return product, err
}

func UpdateProduct(product config.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4`
	_, err := database.PsqlDB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	return err
}

func DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := database.PsqlDB.Exec(query, id)
	return err
}
