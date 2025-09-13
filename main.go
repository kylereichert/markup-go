package main

import (
	"fmt"
	"math"
	"strings"
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

func ConvertToFraction(feet float64) string {
	/*
		Other consideration:
		Should think about splitting this function up and using helper functions
	*/
	precision := 8.0
	feet_floor := math.Floor(feet)
	inch_dec := (feet - feet_floor) * 12
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
			s =
				fmt.Sprintf("%d' %d %d/%d\"", int(feet), int(inch_whole), int(inch_frac), int(precision))
		}
	}

	return s
}

func (i Imperial) AsFraction() string {
	return ConvertToFraction(i.Feet)
}

func ConvertToDecimal(feet string) []string {
	delimiters := "'\" /"

	parts := strings.FieldsFunc(feet, func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	})

	return parts
}

func main() {
	y := Metric{58.7589}
	// x := y.ToImperial()
	// x := Imperial{6.875}
	z := y.ToImperial().AsFraction()
	fmt.Println(z)

	test_frac := "5' 3 7/8\""
	delimiters := " '\"'/"

	parts := strings.FieldsFunc(test_frac, func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	})

	fmt.Println(parts)
}
