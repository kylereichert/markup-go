package main

import (
	"fmt"
	// "fyne.io/fyne/v2/widget"
	"github.com/kylereichert/markup-go/calc"
)

func main() {
	// Some testing for the conversion logic and imperial string formatting
	y := calc.Metric{Meters: 58.7589}
	z := y.ToImperial().AsFraction()

	x := calc.ConvertToDecimal(z)
	r := x.ToMetric()

	fmt.Println(y.Meters)
	fmt.Println(z)
	fmt.Println(x.Feet)
	fmt.Println(r)
	fmt.Println(r.ToImperial().AsFraction())

}
