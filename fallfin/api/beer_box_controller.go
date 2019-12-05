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
	//var currency string
	//var quantity int

	// obtiene info de parameters
	id, _ = strconv.Atoi(c.Param("beerID"))

	//fmt.Print("q onda:  " + strconv.Itoa(id))

	// obtiene info de
	raw, _ := ioutil.ReadAll(c.Request().Body())
	c.Request().SetBody(ioutil.NopCloser(bytes.NewReader(raw)))

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(bytes.NewReader(raw)).Decode(&jsonMap)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request Inválido")
	}

	//json_map has the JSON Payload decoded into a map
	currency, _ := json.Marshal(jsonMap["currency"])
	quantity, _ := json.Marshal(jsonMap["quantity"])

	// busca cerveza por ID
	var db *sql.DB
	db = store.OpenBD()

	var beerItem = store.GetBeerItem(db, id) // printing the user
	if beerItem.Id == 0 {
		return c.String(http.StatusNotFound, "El Id {"+strconv.Itoa(id)+"} de la cerveza no existe")
	}

	var beerBox models.BeerBox
	montoConvertCurrency := service.ConvertCurrency(beerItem.Currency, string(currency), beerItem.Price)
	intQuantity, _ := strconv.Atoi(string(quantity))
	quantityFixBox := service.GetQuantityBeerOK(intQuantity)
	beerBox.PriceTotal = montoConvertCurrency * float32(quantityFixBox)
	beerBox.QuantityFinal = quantityFixBox
	beerBox.BeerBox = quantityFixBox / 6

	fmt.Println(string(currency) + " " + string(quantity))
	return c.JSON(http.StatusOK, beerBox)

}
