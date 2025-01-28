package config

type Order struct {
	ID       int       `json:"ID,omitempty"`
	Buyer    string    `json:"buyer,omitempty"`
	Price    float64   `json:"price,omitempty"`
	Products []Product `json:"products,omitempty"`
}
