package models

type Finance struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
