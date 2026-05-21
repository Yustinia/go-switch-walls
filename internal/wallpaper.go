package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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

func AllocateList() model {
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
