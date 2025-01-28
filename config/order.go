package config

import "time"

type Order struct {
	ID        int       `json:"ID,omitempty"`
	BuyerID   int       `json:"buyerID,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Products  []Product `json:"products,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
