package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"math"
	"strconv"
)

func main() {
	t := strconv.FormatFloat(0, 'f',
		2, 64)

	fmt.Println(t)
}

func FormatNegativePriceWithScores(price float64) string {
	return "−\u00A0" + FormatPrice(price)
}

func FormatPrice(price float64) string {
	rounded := math.Round(price*100) / 100
	withoutFractional := float64(int(rounded))

	p := message.NewPrinter(language.Russian)
	if withoutFractional == rounded { // не отображаем 0 копеек
		return p.Sprint(number.Decimal(withoutFractional))
	}
	return p.Sprint(number.Decimal(rounded, number.IncrementString("0.01")))
}
