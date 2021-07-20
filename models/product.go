package models

type Products struct {
	Id          int     `json:"id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
}

