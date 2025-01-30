package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
	"encoding/json"
	"fmt"
)

// CreateOrder is creating order in next way:
// 1. Get config.Order from request (in request body we don`t have order.Price and order.Products[].Price)
// 2. Order.Price = for range order.Products[] {product.Quantity * product.Price}
// 3. Reduce product.Quantity for each product in order
func CreateOrder(order config.Order) (int, error) {
	var id int
	var sum float64
	var err error
	var productsJSON []byte

	for i := range order.Products {
		product := &order.Products[i]

		productFromDB, err := GetProductByID(product.ID)
		if err != nil {
			fmt.Println("Product not found")
			continue
		}

		sum += float64(product.Quantity) * productFromDB.Price
		product.Price = productFromDB.Price

		productFromDB.Quantity -= product.Quantity
		err = UpdateProduct(productFromDB)
		if err != nil {
			fmt.Println("Error with update")
		}
	}

	productsJSON, err = json.Marshal(order.Products)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO orders (buyer_id, price, products) VALUES ($1, $2, $3) RETURNING id`
	err = database.PsqlDB.QueryRow(query, order.BuyerID, sum, productsJSON).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func GetOrderByID(id int) (config.Order, error) {
	var order config.Order
	var productsJSON []byte
	query := `SELECT id, buyer_id, price, products, created_at FROM orders WHERE id = $1`
	err := database.PsqlDB.QueryRow(query, id).Scan(&order.ID, &order.BuyerID, &order.Price, &productsJSON, &order.CreatedAt)

	err = json.Unmarshal(productsJSON, &order.Products)
	if err != nil {
		return order, err
	}

	return order, err
}

func GetOrders() ([]config.Order, error) {
	var orders []config.Order

	query := `SELECT id, buyer_id, price, products, created_at FROM orders`
	rows, err := database.PsqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order config.Order
		var productsJSON []byte

		err := rows.Scan(&order.ID, &order.BuyerID, &order.Price, &productsJSON, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(productsJSON, &order.Products)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling products JSON: %w", err)
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func UpdateOrder(order config.Order) error {

	var err error
	var productsJSON []byte
	productsJSON, err = json.Marshal(order.Products)

	if err != nil {
		return err
	}

	query := `UPDATE orders SET buyer_id = $1, products = $2 WHERE id = $3`
	_, err = database.PsqlDB.Exec(query, order.BuyerID, productsJSON, order.ID)
	return err
}

func DeleteOrder(id int) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := database.PsqlDB.Exec(query, id)
	return err
}
