package api

import (
  "net/http"
  "github.com/labstack/echo"
  "strconv"
  "database/sql"

  "FallTestMSGo/fallfin/models"
  "FallTestMSGo/fallfin/dblite"
)

/* Buscar Cerveza por ID  */
func SearchBeerByIdGET(c echo.Context) error {
  id,_ := strconv.Atoi(c.Param("beerID"))

  var db *sql.DB
  db = dblite.OpenBD()

  var beerItem = dblite.GetBeerItem(db, id) // printing the user

  // beerIten := new(models.BeerItem)
  //beerIten.Id =id
  if (beerItem.Id != 0){
    return c.JSON(http.StatusOK, beerItem)
  }else{
    return c.JSON(http.StatusNotFound, beerItem)
  }

}

/* Agregar cerveza */
func AddBeers(c echo.Context) error {

  // Bind the input data to ExampleRequest
  /*
  exampleRequest := new(model.ExampleRequest)
  if err := c.Bind(exampleRequest); err != nil {
    return err
  }
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

/* Buscar todas las cerezas */
func SearchBeers(c echo.Context) error {
  //var beer models.BeerItem
	var beers models.BeerItemList
  var db *sql.DB

  db = dblite.OpenBD()
  beers = dblite.SearchAllBeer(db)

  return c.JSON(http.StatusOK, beers)
}
