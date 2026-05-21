package internal

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const pageSize int = 20

type model struct {
	walls   []string
	curPage int
	cursor  int
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

type Styles struct {
	Left  lipgloss.Style
	Right lipgloss.Style
}

func makeStyle() Styles {
	return Styles{
		Left:  lipgloss.NewStyle().Width(80),
		Right: lipgloss.NewStyle().Width(100),
	}
}

func (m model) View() tea.View {
	var b strings.Builder
	styles := makeStyle()
	page := m.currentPage()
	curPageWall := ""

	for index, wall := range page {
		cursor := " "
		if m.cursor == index {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s %s\n", cursor, wall)
	}

	curPageWall = b.String()

	leftCol := styles.Left.Render(curPageWall)
	rightCol := styles.Right.Render("Placeholder")

	render := lipgloss.JoinHorizontal(lipgloss.Left, leftCol, rightCol)
	return tea.NewView(render)
}
