package app

import (
	"ejTestTask/database"
	"ejTestTask/middleware"
	"ejTestTask/transport/json"
	"fmt"
	"net/http"
)

func Run() {
	//Connect and open connect to DB
	database.Connect()
	defer database.PsqlDB.Close()

	//Login endpoint
	http.HandleFunc("/login", json.LoginHandler())

	//All CRUD endpoints for Buyer entity
	http.Handle("/buyer", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			json.CreateBuyerHandler()(w, r)
		case http.MethodGet:
			json.GetBuyerByIDHandler()(w, r)
		case http.MethodPut:
			json.UpdateBuyerHandler()(w, r)
		case http.MethodDelete:
			json.DeleteBuyerHandler()(w, r)
		default:
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})))

	http.Handle("/buyers", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.GetBuyersHandler()(w, r)
	})))

	//All CRUD endpoints for Seller entity
	http.Handle("/seller", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			json.CreateSellerHandler()(w, r)
		case http.MethodGet:
			json.GetSellerByIDHandler()(w, r)
		case http.MethodPut:
			json.UpdateSellerHandler()(w, r)
		case http.MethodDelete:
			json.DeleteSellerHandler()(w, r)
		default:
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})))

	http.Handle("/sellers", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.GetSellersHandler()(w, r)
	})))

	//All CRUD endpoints for Product entity
	http.Handle("/product", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			json.CreateProductHandler()(w, r)
		case http.MethodGet:
			json.GetProductByIDHandler()(w, r)
		case http.MethodPut:
			json.UpdateProductHandler()(w, r)
		case http.MethodDelete:
			json.DeleteProductHandler()(w, r)
		default:
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})))

	http.Handle("/products", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.GetProductsHandler()(w, r)
	})))

	//All CRUD endpoints for Order entity
	http.Handle("/order", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			json.CreateOrderHandler()(w, r)
		case http.MethodGet:
			json.GetOrderByIDHandler()(w, r)
		case http.MethodPut:
			json.UpdateOrderHandler()(w, r)
		case http.MethodDelete:
			json.DeleteOrderHandler()(w, r)
		default:
			http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})))

	http.Handle("/orders", middleware.AuthMiddleware("admin", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.GetOrdersHandler()(w, r)
	})))

	//Starting server
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server. Error^ %v", err)
	}
}
