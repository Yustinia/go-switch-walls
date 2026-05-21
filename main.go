package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
)

const pageSize int = 20

type model struct {
	walls   []string
	curPage int
	cursor  int
}

func applyWallpaper(wall string) error {
	matugenFlags := []string{
		"image", wall, "-t", "scheme-expressive", "-m", "dark", "--contrast", "0.1", "--source-color-index", "0",
	}
	awwwFlags := []string{
		"img", wall, "--transition-type", "simple", "--transition-step", "2", "--transition-fps", "60",
	}

	cmd := exec.Command("matugen", matugenFlags...)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("matugen failed: %w", err)
	}

	cmd = exec.Command("awww", awwwFlags...)

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("awww failed: %w", err)
	}

	return nil
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

func (m model) currentPage() []string {
	start := m.curPage * pageSize
	end := min(start+pageSize, len(m.walls))
	page := m.walls[start:end]

	return page
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var err error
	page := m.currentPage()

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
			if m.curPage > 0 {
				m.cursor = 0
				m.curPage--
			}

		case "right", "l":
			if m.curPage < validPages-1 {
				m.cursor = 0
				m.curPage++
			}

		case "enter":
			selected := m.currentPage()[m.cursor]
			err = applyWallpaper(selected)

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	page := m.currentPage()

	s := ""

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
