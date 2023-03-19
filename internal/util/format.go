package util

import (
	"fmt"
	"os/exec"
)

func FormatCode(fullPathToFile string) error {
	cmd := exec.Command("go", "fmt", fullPathToFile)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error in run go fmt: %w", err)
	}
	return nil
}
