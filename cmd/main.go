package main

import (
	"go-switch-walls/internal"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(internal.AllocateList())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
