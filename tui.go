package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kylereichert/markup-go/calc"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	inputUSJ             = 0
	inputTOJ             = 1
	inputATF             = 2
	inputTOS             = 3
	inputGarageHighPoint = 4
	inputFrontLeft       = 5
	inputFrontRight      = 6
	inputBackLeft        = 7
	inputBackRight       = 8
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("150"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("150"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	errorStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func initialModel() model {
	const inputWidth = 20
	const defaultInputLim = 10
	m := model{
		inputs: make([]textinput.Model, 9),
	}

	numericValidator := func(s string) error {
		if s == "" {
			return nil
		}
		if _, err := strconv.ParseFloat(s, 64); err != nil {
			return fmt.Errorf("Must be a number")
		}
		return nil
	}

	// var t textinput.Model
	for i := range m.inputs {
		t := textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = defaultInputLim
		t.Width = inputWidth
		t.Validate = numericValidator

		switch i {
		case inputUSJ:
			t.Placeholder = "Enter USJ"
			t.Width = inputWidth
			t.Focus()
			// t.CharLimit = defaultInputLim
			// t.TextStyle = focusedStyle
			// t.PromptStyle = focusedStyle
		case inputTOJ:
			t.Placeholder = "Enter TOJ"
		case inputATF:
			t.Placeholder = "Enter ATF"
		case inputTOS:
			t.Placeholder = "Garage TOS"
		case inputGarageHighPoint:
			t.Placeholder = "Garage High Point"
		case inputFrontLeft:
			t.Placeholder = "Front Left"
		case inputFrontRight:
			t.Placeholder = "Front Right"
		case inputBackLeft:
			t.Placeholder = "Back Left"
		case inputBackRight:
			t.Placeholder = "Back Right"
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			// toj, _ := strconv.ParseFloat(m.inputs[inputTOJ].Value(), 64)
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		// b.WriteString(m.inputs[i].View())
		view := m.inputs[i].View()

		// Testing input handling. Im thinking a switch statement would work
		// better than a loop here
		if m.inputs[i].Value() != "" && m.inputs[i].Err == nil {
			val, _ := strconv.ParseFloat(m.inputs[i].Value(), 64)
			// Covers necessary use cases, but should consider arbitrary calcs
			switch i {
			case inputUSJ:
				// metricUSJ := calc.Metric{Meters: val}
				// imperialUSJ := metricUSJ.ToImperial().AsFraction()
				view += fmt.Sprintf("USJ: %.2f", val)
			case inputTOJ:
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				beamSize := calc.Metric{Meters: val - usj}.ToImperial().AsFraction()
				view += fmt.Sprintf("Beam Size: %s", beamSize)
			case inputATF:
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				wallHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Wall Height: %s", wallHeight)
			case inputTOS:
				// Consider adding that 5" maybe? Or should I keep doing it mentally?
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				slabHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Garage Opening Drop: %s", slabHeight)
			case inputGarageHighPoint:
				// Again, should i account for the 9" or no?
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				garageWallHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Garage Wall Drop: %s", garageWallHeight)
			case inputFrontLeft:
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				frontLeftHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Front Left Drop: %s", frontLeftHeight)
			case inputFrontRight:
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				frontRightHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Front Right Drop: %s", frontRightHeight)
			case inputBackLeft:
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				backLeftHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Back Left Drop: %s", backLeftHeight)
			case inputBackRight:
				usj, _ := strconv.ParseFloat(m.inputs[inputUSJ].Value(), 64)
				backRightHeight := calc.Metric{Meters: usj - val}.ToImperial().AsFraction()
				view += fmt.Sprintf("Back Right Drop: %s", backRightHeight)
			}
		}

		if m.inputs[i].Err != nil {
			view = errorStyle.Render(view)
			view += m.inputs[i].Err.Error()
		}

		b.WriteString(view)
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
