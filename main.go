package main

import (
	"fmt"
	"os"
	// "github.com/kylereichert/markup-go/calc"

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

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Could not start program: %s\n", err)
		os.Exit(1)
	}
}
