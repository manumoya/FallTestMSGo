package models

type BeerItem struct {
  Id         int `json:"id"`
  Name       string `json:"name"`
  Brewery    string `json:"brewery"`
  //I created a struct with a struct to select the rows in the table and add data.
}
