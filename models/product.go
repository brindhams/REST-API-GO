package models

type product struct {
	Id          string  `json:"id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
}

