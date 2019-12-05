package models

type BeerBox struct {
	PriceTotal    float32 `json:"priceTotal"`
	QuantityFinal int     `json:"quantityFinal"`
	BeerBox       int     `json:"beerBox"`
}

/*
type BeerBoxInput struct {
  Currency  int `json:"currency"`
  Quantity  int `json:"quantity"`
}
*/
