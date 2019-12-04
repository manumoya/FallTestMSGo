package api

import (
  "FallTestMSGo/fallfin/models"
  "net/http"
  "github.com/labstack/echo"
  "strconv"
  "database/sql"
  "FallTestMSGo/fallfin/dblite"
)


func SearchBeerByIdGET(c echo.Context) error {

  id,_ := strconv.Atoi(c.Param("id"))
  name,_ := strconv.Atoi(c.Param("name"))
  brewery,_ := strconv.Atoi(c.Param("brewery"))

  beerIten := new(models.BeerItem)
  beerIten.Id =id
  beerIten.Name=string(name)
  beerIten.Brewery=string(brewery)

  return c.JSON(http.StatusOK, beerIten)

}


/* Agregar cerveza */
func AddBeers(c echo.Context) error {

  // Bind the input data to ExampleRequest
  /*
  exampleRequest := new(model.ExampleRequest)
  if err := c.Bind(exampleRequest); err != nil {
    return err
  }

  // Manipulate the input data
  greeting := exampleRequest.FirstName + " " + exampleRequest.LastName

  return c.JSONBlob(
    http.StatusOK,
    []byte(
      fmt.Sprintf(`{
        "first_name": %q,
        "last_name": %q,
        "msg": "Hello %s"
      }`, exampleRequest.FirstName, exampleRequest.LastName, greeting),
    ),
  )
  */

  beerIten := new(models.BeerItem)
  if err := c.Bind(beerIten); err != nil {
    return err
  }

  var db *sql.DB
  db = dblite.OpenBD()
  dblite.AddBeerItem(db, beerIten.Id, beerIten.Name, beerIten.Brewery,
                     beerIten.Country, beerIten.Price, beerIten.Currency)

  return c.JSON(http.StatusOK, beerIten)
}
