package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type MyMode struct {
	quitting  bool
	altscreen bool
}

var _ tea.Model = &MyMode{}

func (m *MyMode) Init() tea.Cmd {
	return nil
}

func (m *MyMode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch in := msg.(type) {
	case tea.KeyMsg:
		switch in.String() {
		case "q", "ctrl+c", "ctrl+w", "esc":
			m.quitting = true
			return m, tea.Quit
		case " ", "tab":
			m.altscreen = !m.altscreen
			return m, tea.EnterAltScreen
		}
	}
	return m, nil
}

func (m *MyMode) View() string {
	if m.quitting {
		return "Good Bye! 🐍"
	}
	var diff string
	if m.altscreen {
		diff = "here is A tab! boy"
	} else {
		diff = "now you are standing in B"
	}
	return fmt.Sprintf("hello! \r\n \r\n %s \r\n you can use space to allscreen tab or Q to quit \r\n", diff)
}

func main() {
	program := tea.NewProgram(&MyMode{})
	_, err := program.Run()
	if err != nil {
		panic("run program err")
	}
}
