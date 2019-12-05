package models

type BeerBox struct {
  Price_total   float32 `json:"price_total"`
}

type BeerBoxInput struct {
  Currency  int `json:"currency"`
  Quantity  int `json:"quantity"`
}
