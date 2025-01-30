package json

import (
	"ejTestTask/config"
	"ejTestTask/database/postgresql"
	"ejTestTask/middleware"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error":"Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		var user config.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON payload. Error: %v"}`, err), http.StatusBadRequest)
			return
		}

		user, err := postgresql.GetUserRole(user.Login, user.Password)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Failed to get user. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		token, err := middleware.GenerateJWT(user.Login, user.Role)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"Could not generate token. Error: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"Token": token})
	}
}
