package service

/*ConvertCurrency => Convierte moneda*/
func ConvertCurrency(currBeerCode string, currBuyCode string, monto float32) float32 {
	var usd float32
	//var clp float32
	usd = 819.0
	//clp = 1.0

	/*
		fmt.Println("com  " + currBeerCode + " / " + currBuyCode)
		fmt.Println(reflect.TypeOf(currBeerCode))
		fmt.Println(reflect.TypeOf(currBuyCode))
	*/

	if currBeerCode == "USD" && currBuyCode == "CLP" {
		//fmt.Println("conv1")
		return monto * usd
	} else if currBeerCode == "CLP" && currBuyCode == "USD" {
		//fmt.Println("conv2")
		return monto / usd
	} else {
		//fmt.Println("conv3")
		return monto
	}
}
