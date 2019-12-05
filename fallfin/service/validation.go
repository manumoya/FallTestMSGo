package service

func GetQuantityBeerOK(quantity int) int {
	mod6 := quantity % 6
	if mod6 > 0 {
		return (quantity + (6 - mod6))
	}
	return quantity

}
