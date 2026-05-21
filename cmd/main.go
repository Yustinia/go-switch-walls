package main

import (
	"go-switch-walls/internal"
	"os"

	tea "charm.land/bubbletea/v2"
)

func checkInstalledPkgs() error {
	pkgs := []string{"matugen", "awww"}

	for _, pkg := range pkgs {
		err := internal.CommandExists(pkg)

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := checkInstalledPkgs()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(internal.AllocateList())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
