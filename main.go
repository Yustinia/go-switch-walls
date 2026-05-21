package main

import (
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	walls []string
}

func allocateList() model {
	return model{walls: []string{"Test1", "Test2"}}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() tea.View {
	msg := m.walls[0]
	return tea.NewView(msg)
}

func main() {
	p := tea.NewProgram(allocateList())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
