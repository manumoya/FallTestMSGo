package models

type BeerItem struct {
  Id        int `json:"id"`
  Name      string `json:"name"`
  Brewery   string `json:"brewery"`
  Country   string `json:"country"`
  Price     float32 `json:"price"`
  Currency  string `json:"currency"`
}
