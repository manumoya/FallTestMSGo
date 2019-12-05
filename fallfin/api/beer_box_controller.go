package api

import (
  "strconv"
  "github.com/labstack/echo"
  "net/http"
  "fmt"
)

func BoxBeerPriceById(c echo.Context) error {

  id,_ := strconv.Atoi(c.Param("beerID"))
  currency_buy := c.Param("currency")
  quantity_buy,_ := strconv.Atoi(c.Param("quantity"))

  fmt.Println("currency_buy: " + currency_buy)

  var msg= "Id {"+ strconv.Itoa(id) +"} / currency_buy" + currency_buy + " / quantity_buy "+ strconv.Itoa(quantity_buy)

  return c.String(http.StatusOK, msg)
}
