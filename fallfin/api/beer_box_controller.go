package api

import (
	"bytes"
	"encoding/json"
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
		return c.String(http.StatusBadRequest, "malo malo ")
	}

	//json_map has the JSON Payload decoded into a map
	currency, _ := json.Marshal(jsonMap["currency"])
	quantity, _ := json.Marshal(jsonMap["quantity"])

	return c.String(http.StatusOK, "r: "+strconv.Itoa(id)+" "+string(currency)+" "+string(quantity))

}
