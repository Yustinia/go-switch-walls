package internal

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const pageSize int = 20

type state int
type colorMode int

const (
	dark colorMode = iota
	light
)
const (
	leftFoc state = iota
	rightFoc
)

var matugenSchemes = []string{
	"scheme-content", "scheme-expressive", "scheme-fidelity", "scheme-fruit-salad", "scheme-monochrome", "scheme-neutral", "scheme-rainbow", "scheme-tonal-spot", "scheme-vibrant",
}

type model struct {
	walls        []string
	curPage      int
	wallCursor   int
	schemeCursor int
	curState     state
	color        colorMode
	width        int
	height       int
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
			if m.curState == leftFoc {
				if m.wallCursor > 0 {
					m.wallCursor--
				}
			} else if m.curState == rightFoc {
				if m.schemeCursor > 0 {
					m.schemeCursor--
				}
			}

		case "down", "j":
			if m.curState == leftFoc {
				if m.wallCursor < len(page)-1 {
					m.wallCursor++
				}
			} else if m.curState == rightFoc {
				if m.schemeCursor < len(matugenSchemes)-1 {
					m.schemeCursor++
				}
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
			selWall := m.currentPage()[m.wallCursor]
			selScheme := matugenSchemes[m.schemeCursor]
			selMode := m.setColorMode()
			err = applyWallpaper(selWall, selScheme, selMode)

			if err != nil {
				fmt.Println(err)
			}

		case "tab":
			m.curState = (m.curState + 1) % 2

		case "s":
			if m.color == dark {
				m.color = light
			} else {
				m.color = dark
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

type Styles struct {
	Left  lipgloss.Style
	Right lipgloss.Style

	RightTopRow lipgloss.Style
	RightBotRow lipgloss.Style
}

func makeStyle(totalWidth int, totalHeight int) Styles {
	leftWidth := 80
	rightWidth := totalWidth - leftWidth

	rightBotRowHeight := 4
	rightTopRowHeight := totalHeight - rightBotRowHeight

	return Styles{
		Left:        lipgloss.NewStyle().Width(leftWidth).Height(totalHeight).Border(lipgloss.RoundedBorder()).AlignVertical(lipgloss.Center),
		Right:       lipgloss.NewStyle().Width(rightWidth),
		RightTopRow: lipgloss.NewStyle().Width(rightWidth).Height(rightTopRowHeight).Border(lipgloss.RoundedBorder()).AlignVertical(lipgloss.Center),
		RightBotRow: lipgloss.NewStyle().Width(rightWidth).Height(rightBotRowHeight).Border(lipgloss.RoundedBorder()).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center),
	}
}

func (m model) makeWallList() string {
	var b strings.Builder
	page := m.currentPage()

	for index, wall := range page {
		cursor := " "
		if m.wallCursor == index {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s %s\n", cursor, wall)
	}
	return b.String()
}

func (m model) makeSchemeList() string {
	var b strings.Builder

	for index, scheme := range matugenSchemes {
		cursor := " "
		if m.schemeCursor == index {
			cursor = ">"
		}

		fmt.Fprintf(&b, "%s %s\n", cursor, scheme)
	}
	return b.String()
}

func (m model) setColorMode() string {
	colorStr := "dark"
	if m.color == dark {
		colorStr = "dark"
	} else {
		colorStr = "light"
	}

	return colorStr
}

func viewColorMode(colorStr string) string {
	return fmt.Sprintf("[S] Mode: %s", colorStr)
}

func (m model) View() tea.View {
	styles := makeStyle(m.width, m.height)
	wallList := m.makeWallList()
	schemeList := m.makeSchemeList()
	colorState := m.setColorMode()
	colorStr := viewColorMode(colorState)

	leftCol := styles.Left.Render(wallList)
	rightCol := styles.Right.Render(lipgloss.JoinVertical(lipgloss.Left, styles.RightTopRow.Render(schemeList), styles.RightBotRow.Render(colorStr)))

	render := lipgloss.JoinHorizontal(lipgloss.Left, leftCol, rightCol)
	viewer := tea.NewView(render)
	viewer.AltScreen = true

	return viewer
}
