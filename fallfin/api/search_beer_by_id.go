package api

import (

  "FallTestMSGo/fallfin/models"

  "fmt"
  "net/http"

  "github.com/labstack/echo"

  "strconv"

  //"gitlab.com/ykyuen/golang-echo-template-example/model"

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



func SearchBeerById(c echo.Context) error {

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

  fmt.Printf("Buscar ini \n")

  beerIten := new(models.BeerItem)
  if err := c.Bind(beerIten); err != nil {
    return err
  }

  fmt.Printf("Beer Name" + beerIten.Name +" \n")

  fmt.Printf("Buscar fin \n")


  /*
  return c.JSONBlob(
    http.StatusOK,
    []byte(
      fmt.Sprintf(`{
        "id": %q,
        "name": %q,
        "brewery": %q
      }`, beerIten.Id, beerIten.Name, beerIten.Brewery ),
    ),
  )
  */

  return c.JSON(http.StatusOK, beerIten)



}
