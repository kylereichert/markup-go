package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)


const (
	inputTOJ = 0
	inputUSJ = 1
	inputATF = 2
	inputTOS = 3
	inputGarageHighPoint = 4
	inputFrontLeft = 5
	inputFrontRight = 6
	inputBackLeft = 7
	inputBackRight = 8
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
		case inputTOJ:
			t.Placeholder = "Enter TOJ"
			t.Width = inputWidth
			t.Focus()
			// t.CharLimit = defaultInputLim
			// t.TextStyle = focusedStyle
			// t.PromptStyle = focusedStyle
		case inputUSJ:
			t.Placeholder = "Enter USJ"
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
		case "ctrl+c", "esc":
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

			switch i {
			case inputUSJ:
				view += fmt.Sprintf("USJ: %.2f", val)
			case inputTOJ:
				view += fmt.Sprintf("TOJ: %.2f", val)
			case inputATF:
				view += fmt.Sprintf("ATF: %.2f", val)
			case inputTOS:
				view += fmt.Sprintf("TOS: %.2f", val)
			case inputGarageHighPoint:
				view += fmt.Sprintf("TBD: %.2f", val)
			case inputFrontLeft:
				view += fmt.Sprintf("TBD: %.2f", val)
			case inputFrontRight:
				view += fmt.Sprintf("TBD: %.2f", val)
			case inputBackLeft:
				view += fmt.Sprintf("TBD: %.2f", val)
			case inputBackRight:
				view += fmt.Sprintf("TBD: %.2f", val)
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

