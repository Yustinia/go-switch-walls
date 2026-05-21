package internal

import (
	"fmt"
	"os/exec"
)

func CommandExists(pkg string) error {
	_, err := exec.LookPath(pkg)

	if err != nil {
		return fmt.Errorf("%s does not exist: %w", pkg, err)
	}

	return nil
}
