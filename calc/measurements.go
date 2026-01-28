package calc

import (
	"fmt"
	"math"
	"strconv"
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
		TODO: Does not handle negative inputs, AFAIK
	*/
	precision := 8.0
	feet_floor := math.Floor(feet)
	inch_dec := (feet - feet_floor) * 12
	inch_whole := math.Floor(inch_dec)
	inch_frac := math.Round((inch_dec - inch_whole) * precision)
	isNegative := false

	if feet < 0 {
		isNegative = true
	}

	// Truncate the fraction if it is zero. Else, reduce fraction.
	var s string
	if int(inch_frac) == 0 {
		switch isNegative {
		case true:
			s = fmt.Sprintf("-%d' %d\"", int(feet), int(inch_whole))
		case false:
			s = fmt.Sprintf("%d' %d\"", int(feet), int(inch_whole))
		}
	} else {
		for int(inch_frac)%2 == 0 {
			inch_frac = inch_frac / 2
			precision = precision / 2
		}
		// in case the fractional component rounds to a whole number

		switch isNegative {
		case true:
			if inch_frac == precision {
				inch_whole += 1
				s = fmt.Sprintf("%d' %d\"", int(feet), int(inch_whole))
			} else {
				s =
					fmt.Sprintf("%d' %d %d/%d\"", int(feet), int(inch_whole), int(inch_frac), int(precision))
			}
		case false:
			if inch_frac == precision {
				inch_whole += 1
				s = fmt.Sprintf("%d' %d\"", int(feet), int(inch_whole))
			} else {
				s =
					fmt.Sprintf("%d' %d %d/%d\"", int(feet), int(inch_whole), int(inch_frac), int(precision))
			}
		}
	}

	return s
}

func (i Imperial) AsFraction() string {
	return ConvertToFraction(i.Feet)
}

func ConvertToDecimal(feet string) Imperial {
	// Currently needs a footage or it will panic. i.e. 5" does not work, so use 0' 5"
	// Should fix this in the future
	delimiters := "'\" /"

	strParts := strings.FieldsFunc(feet, func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	})

	intParts := make([]int, len(strParts))

	// Convert the string array into a int array
	for i, s := range strParts {
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("error converting:", s, err)
		}
		intParts[i] = val
	}

	var decimalInch float64
	var decimalFoot float64

	partsLen := len(intParts)

	if partsLen < 4 {
		decimalInch = float64(intParts[1]) / 12.0
	} else {
		decimalInch = (float64(intParts[1]) + float64(intParts[2])/float64(intParts[3])) / 12.0
	}

	decimalFoot = float64(intParts[0]) + decimalInch

	return Imperial{Feet: decimalFoot}
}
