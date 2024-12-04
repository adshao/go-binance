package common

import (
	"bytes"
	"fmt"
	"math"
)

// AmountToLotSize converts an amount to a lot sized amount
func AmountToLotSize(lot float64, precision int, amount float64) float64 {
	return math.Trunc(math.Floor(amount/lot)*lot*math.Pow10(precision)) / math.Pow10(precision)
}

func RoundPriceToTickSize(price, tickSize float64) float64 {
	if tickSize == 0 {
		return price // To avoid division by zero
	}
	// Calculate the factor by which to multiply and divide the price to conform to the tick size.
	factor := 1 / tickSize

	// Round the price to the nearest tick size.
	roundedPrice := math.Round(price*factor) / factor

	return roundedPrice
}

// ToJSONList convert v to json list if v is a map
func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}

func ToInt(digit interface{}) (i int, err error) {
	if intVal, ok := digit.(int); ok {
		return int(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}

func ToInt64(digit interface{}) (i int64, err error) {
	if intVal, ok := digit.(int); ok {
		return int64(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int64(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}
