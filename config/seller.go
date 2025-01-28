package config

type Seller struct {
	ID    int    `json:"ID,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}
