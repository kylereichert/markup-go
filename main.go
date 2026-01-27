package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)
func main() {
	// Some testing for the conversion logic and imperial string formatting
	// y := calc.Metric{Meters: 58.7589}
	// z := y.ToImperial().AsFraction()

	// x := calc.ConvertToDecimal(z)
	// r := x.ToMetric()

	// fmt.Println(y.Meters)
	// fmt.Println(z)
	// fmt.Println(x.Feet)
	// fmt.Println(r)
	// fmt.Println(r.ToImperial().AsFraction())

	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("Could not start program: %s\n", err)
		os.Exit(1)
	}
}
