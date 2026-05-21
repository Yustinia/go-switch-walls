package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
)

const pageSize int = 20

type model struct {
	walls  []string
	page   int
	cursor int
}

func allocateList() model {
	var wallList []string

	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	wallPath := filepath.Join(homePath, "Pictures", "Walls")

	entries, err := os.ReadDir(wallPath)
	if err != nil {
		panic(err)
	}

	for _, wall := range entries {
		fullPath := filepath.Join(wallPath, wall.Name())
		wallList = append(wallList, fullPath)
	}

	return model{walls: wallList}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	start := m.page * pageSize
	end := min(start+pageSize, len(m.walls))
	page := m.walls[start:end]

	validPages := len(m.walls) / pageSize
	if len(m.walls)%pageSize != 0 {
		validPages++
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(page)-1 {
				m.cursor++
			}
		case "left", "h":
			if m.page > 0 {
				m.cursor = 0
				m.page--
			}
		case "right", "l":
			if m.page < validPages-1 {
				m.cursor = 0
				m.page++
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := ""
	start := m.page * pageSize
	end := min(start+pageSize, len(m.walls))
	page := m.walls[start:end]

	for index, wall := range page {
		cursor := " "
		if m.cursor == index {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, wall)
	}
	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(allocateList())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
