package cli

import (
	"fmt"
	"genos/internal/service"
	"os/exec"
)

type ExecuteCLI struct {
}

func NewExecuteSLI() *ExecuteCLI {
	return &ExecuteCLI{}
}

var _ service.ExecuteCLI = (*ExecuteCLI)(nil)

func (e ExecuteCLI) ExecuteCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error in run go fmt: %w", err)
	}
	return nil
}
