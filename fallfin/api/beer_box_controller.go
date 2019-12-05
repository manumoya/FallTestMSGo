package api

import (
	"FallTestMSGo/fallfin/models"
	"FallTestMSGo/fallfin/service"
	"FallTestMSGo/fallfin/store"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

/*BoxBeerPriceByID => Lista el precio de una caja de cervezas de una marca*/
func BoxBeerPriceByID(c echo.Context) error {

	var id int

	// obtiene info de parameters
	id, _ = strconv.Atoi(c.Param("beerID"))

	// obtiene info de
	raw, _ := ioutil.ReadAll(c.Request().Body())
	c.Request().SetBody(ioutil.NopCloser(bytes.NewReader(raw)))

	// OJO
	var beerBoxInput models.BeerBoxInput

	err := json.Unmarshal(raw, &beerBoxInput)
	if err != nil {
		//panic(err)
	}

	// busca cerveza por ID
	var db *sql.DB
	db = store.OpenBD()

	var beerItem = store.GetBeerItem(db, id) // printing the user
	if beerItem.Id == 0 {
		return c.String(http.StatusNotFound, "El Id {"+strconv.Itoa(id)+"} de la cerveza no existe")
	}

	var beerBox models.BeerBox

	fmt.Print("conv: " + beerItem.Currency + "  " + beerBoxInput.Currency)
	montoConvertCurrency := service.ConvertCurrency(string(beerItem.Currency), beerBoxInput.Currency, beerItem.Price)

	s := strconv.FormatFloat(float64(montoConvertCurrency), 'f', -1, 32)
	fmt.Print("monto cnvertido: " + s + "\n")

	quantityFixBox := service.GetQuantityBeerOK(beerBoxInput.Quantity)
	beerBox.PriceTotal = montoConvertCurrency * float32(quantityFixBox)
	beerBox.QuantityFinal = quantityFixBox
	beerBox.BeerBox = quantityFixBox / 6

	return c.JSON(http.StatusOK, beerBox)
}
