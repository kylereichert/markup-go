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

func ConvertToFraction(feet float64, precision ...float64) string {
	/*
		Maybe split into helper functions. Using Modf here would be a better way
		to split the whole and fractional/decimal part
	*/
	actualPrecision := 8.0
	if precision != nil {
		actualPrecision = precision[0]
	}

	feet_floor, inch_dec := math.Modf(feet)
	inch_whole, inch_frac := math.Modf(math.Abs(inch_dec * 12)) // Convert to inches
	inch_frac = math.Round(inch_frac * actualPrecision)
	feet_floor = math.Abs(feet_floor)

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
			actualPrecision = actualPrecision / 2
		}

		/*
			Might need to cascade the fractional inch to the whole inch, and then
			possibly the footage. I think the only place where the rounding will
			have to occur will be when the numerator rounds up to match the precision.
			A greater numerator should not be possible
		*/

		if inch_frac == actualPrecision {
			inch_frac, actualPrecision = 0, 0
			inch_whole++
			if inch_whole >= 12 {
				inch_whole = float64(int(inch_whole) - (int(inch_whole) - 12))
				feet_floor++
			}
		}

		// in case the fractional component rounds to a whole number

		switch isNegative {
		case true:
			if inch_frac == 0 {
				s = fmt.Sprintf("%d' %d\"", int(feet_floor), int(inch_whole))
			} else {
				s = fmt.Sprintf("%d' %d %d/%d\"", int(feet), int(inch_whole), int(inch_frac), int(actualPrecision))
			}
		case false:
			if inch_frac == 0 {
				s = fmt.Sprintf("%d' %d\"", int(feet), int(inch_whole))
			} else {
				s = fmt.Sprintf("%d' %d %d/%d\"", int(feet), int(inch_whole), int(inch_frac), int(actualPrecision))
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
	delimiters := "' \" /"

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
