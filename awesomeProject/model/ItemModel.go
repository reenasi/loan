package model

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
