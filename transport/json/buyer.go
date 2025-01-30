package json

import (
	"database/sql"
	"ejTestTask/config"
	"ejTestTask/database/postgresql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateBuyerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var buyer config.Buyer
		if err := json.NewDecoder(r.Body).Decode(&buyer); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		id, err := postgresql.CreateBuyer(buyer)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to create buyer. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

func GetBuyerByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var buyer config.Buyer

		if err := json.NewDecoder(r.Body).Decode(&buyer); err != nil || buyer.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		buyer, err := postgresql.GetBuyerByID(buyer.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Buyer not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch buyer. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(buyer)
	}
}

func GetBuyersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		buyers, err := postgresql.GetBuyers()
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"error":"Order not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf(`{"error":"Failed to fetch buyers. Error: %v"}`, err), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(buyers)
	}
}

func UpdateBuyerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var buyer config.Buyer

		if err := json.NewDecoder(r.Body).Decode(&buyer); err != nil || buyer.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		if err := postgresql.UpdateBuyer(buyer); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to update buyer. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Buyer %d updated", buyer.ID)})
	}
}

func DeleteBuyerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var buyer config.Buyer
		if err := json.NewDecoder(r.Body).Decode(&buyer); err != nil || buyer.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		if err := postgresql.DeleteBuyer(buyer.ID); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to delete buyer. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Buyer %d deleted", buyer.ID)})
	}
}
