package internal

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const pageSize int = 20

var matugenSchemes = []string{
	"scheme-content", "scheme-expressive", "scheme-fidelity", "scheme-fruit-salad", "scheme-monochrome", "scheme-neutral", "scheme-rainbow", "scheme-tonal-spot", "scheme-vibrant",
}

type model struct {
	walls        []string
	curPage      int
	wallCursor   int
	schemeCursor int
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
			if m.wallCursor > 0 {
				m.wallCursor--
			}

		case "down", "j":
			if m.wallCursor < len(page)-1 {
				m.wallCursor++
			}

		case "left", "h":
			if m.curPage > 0 {
				m.wallCursor = 0
				m.curPage--
			}

		case "right", "l":
			if m.curPage < validPages-1 {
				m.wallCursor = 0
				m.curPage++
			}

		case "enter":
			selected := m.currentPage()[m.wallCursor]
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

	for index, wall := range page {
		cursor := " "
		if m.wallCursor == index {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s %s\n", cursor, wall)
	}
	curPageWall := b.String()
	b.Reset()

	for index, scheme := range matugenSchemes {
		cursor := " "
		if m.schemeCursor == index {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s %s\n", cursor, scheme)
	}
	schemeSel := b.String()
	b.Reset()

	leftCol := styles.Left.Render(curPageWall)
	rightCol := styles.Right.Render(schemeSel)

	render := lipgloss.JoinHorizontal(lipgloss.Left, leftCol, rightCol)
	return tea.NewView(render)
}
