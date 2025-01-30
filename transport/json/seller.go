package json

import (
	"database/sql"
	"ejTestTask/config"
	"ejTestTask/database/postgresql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateSellerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var seller config.Seller
		if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		id, err := postgresql.CreateSeller(seller)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to create seller. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

func GetSellerByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var seller config.Seller

		if err := json.NewDecoder(r.Body).Decode(&seller); err != nil || seller.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		seller, err := postgresql.GetSellerByID(seller.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Seller not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch seller. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(seller)
	}
}

func GetSellersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		sellers, err := postgresql.GetSellers()
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Order not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch sellers. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sellers)
	}
}

func UpdateSellerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var seller config.Seller

		if err := json.NewDecoder(r.Body).Decode(&seller); err != nil || seller.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		if err := postgresql.UpdateSeller(seller); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to update seller. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Seller %d updated", seller.ID)})
	}
}

func DeleteSellerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var seller config.Seller
		if err := json.NewDecoder(r.Body).Decode(&seller); err != nil || seller.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		if err := postgresql.DeleteSeller(seller.ID); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to delete seller. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Seller %d deleted", seller.ID)})
	}
}
