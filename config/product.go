package config

type Product struct {
	ID          int     `json:"ID,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Seller      string  `json:"seller,omitempty"`
}
