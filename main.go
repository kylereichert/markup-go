package main

import (
	"fmt"
	"math"
)

type Imperial struct {
	Feet float64
}

type Metric struct {
	Meters float64
}

func (m Metric) ToImperial() Imperial {
	return Imperial{
		Feet: m.Meters * 3.28084,
	}
}

func (i Imperial) ToMetric() Metric {
	return Metric{
		Meters: i.Feet * 0.3048,
	}
}

func (i Imperial) ToString() string {
	/*
		Other consideration:
		Should think about splitting this function up and using helper functions
	*/
	precision := 8.0
	feet := math.Floor(i.Feet)
	inch_dec := (i.Feet - feet) * 12
	inch_whole := math.Floor(inch_dec)
	inch_frac := math.Round((inch_dec - inch_whole) * precision)

	// Truncate the fraction if it is zero. Else, reduce fraction.
	var s string
	if int(inch_frac) == 0 {
		s = fmt.Sprintf("%d' %d\"", int(feet), int(inch_whole))
	} else {
		for int(inch_frac)%2 == 0 {
			inch_frac = inch_frac / 2
			precision = precision / 2
		}
		// in case the fractional component rounds to a whole number
		if inch_frac == precision {
			inch_whole += 1
			s = fmt.Sprintf("%d' %d\"", int(feet), int(inch_whole))
		} else {
			s = fmt.Sprintf("%d' %d %d/%d\"", int(feet), int(inch_whole), int(inch_frac), int(precision))
		}
	}

	return s
}

func main() {
	y := Metric{58.7589}
	x := y.ToImperial()
	// x := Imperial{6.875}
	z := x.ToString()
	fmt.Println(z)

}
