package config

type Product struct {
	ID          int     `json:"ID,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	SellerID    int     `json:"sellerID,omitempty"`
}
