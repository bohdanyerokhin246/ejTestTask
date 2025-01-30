package json

import (
	"database/sql"
	"ejTestTask/config"
	"ejTestTask/database/postgresql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var order config.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, `{"error":"Invalid JSON payload"}`, http.StatusBadRequest)
			return
		}

		id, err := postgresql.CreateOrder(order)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":  "Failed to create order. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

func GetOrderByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var order config.Order

		if err := json.NewDecoder(r.Body).Decode(&order); err != nil || order.ID <= 0 {
			http.Error(w, `{"error":"Invalid JSON payload"}`, http.StatusBadRequest)
			return
		}

		order, err := postgresql.GetOrderByID(order.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Order not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch order. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}
}

func GetOrdersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		orders, err := postgresql.GetOrders()
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Order not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch order. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}

func UpdateOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var order config.Order

		if err := json.NewDecoder(r.Body).Decode(&order); err != nil || order.ID <= 0 {
			http.Error(w, `{"error":"Invalid JSON payload"}`, http.StatusBadRequest)
			return
		}

		if err := postgresql.UpdateOrder(order); err != nil {
			http.Error(w, `{"error":"Failed to update order"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Order updated"})
	}
}

func DeleteOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var order config.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil || order.ID <= 0 {
			http.Error(w, `{"error":"Invalid JSON payload"}`, http.StatusBadRequest)
			return
		}

		if err := postgresql.DeleteOrder(order.ID); err != nil {
			http.Error(w, `{"error":"Failed to delete order"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Order %d deleted", order.ID)})
	}
}
