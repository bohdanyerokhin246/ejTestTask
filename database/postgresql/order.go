package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func CreateOrder(order config.Order) (int, error) {
	var id int
	query := `INSERT INTO orders (buyer_id, price, products) VALUES ($1, $2, $3) RETURNING id`
	err := database.PsqlDB.QueryRow(query, order.BuyerID, order.Price, order.Products).Scan(&id)
	return id, err
}

func GetOrderByID(id int) (config.Order, error) {
	var order config.Order
	query := `SELECT id, buyer_id, price, products, created_at FROM orders WHERE id = $1`
	err := database.PsqlDB.QueryRow(query, id).Scan(&order.ID, &order.BuyerID, &order.Price, &order.Products, &order.CreatedAt)
	return order, err
}

func UpdateOrder(order config.Order) error {
	query := `UPDATE orders SET buyer_id = $1, products = $2 WHERE id = $3`
	_, err := database.PsqlDB.Exec(query, order.BuyerID, order.Products, order.ID)
	return err
}

func DeleteOrder(id int) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := database.PsqlDB.Exec(query, id)
	return err
}
