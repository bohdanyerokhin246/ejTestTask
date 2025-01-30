package json

import (
	"database/sql"
	"ejTestTask/config"
	"ejTestTask/database/postgresql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var product config.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		id, err := postgresql.CreateProduct(product)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to create order. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

func GetProductByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var product config.Product

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil || product.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		product, err := postgresql.GetProductByID(product.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch order. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

func GetProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		products, err := postgresql.GetProducts()
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Order not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch order. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func UpdateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var product config.Product

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil || product.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		if err := postgresql.UpdateProduct(product); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to update product. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Product %d updated", product.ID)})
	}
}

func DeleteProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var product config.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil || product.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		if err := postgresql.DeleteProduct(product.ID); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to delete product. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Product %d deleted", product.ID)})
	}
}
