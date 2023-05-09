package service

import (
	"fmt"
)

type CliUC struct {
	exec ExecuteCLI
}

func NewCliUC(exec ExecuteCLI) *CliUC {
	return &CliUC{exec: exec}
}

var _ CliCommandContract = (*CliUC)(nil)

func (c *CliUC) Format(fullPathToFile string) error {
	err := c.exec.ExecuteCommand("go", []string{"fmt", fullPathToFile}...)
	if err != nil {
		return fmt.Errorf("error Format: %w", err)
	}
	return nil
}

func (c *CliUC) CreateGoModule(nameProject string) error {
	err := c.exec.ExecuteCommand("go", []string{"mod", "init", nameProject}...)
	if err != nil {
		return fmt.Errorf("error CreateGoModule: %w", err)
	}
	return nil
}

func (c *CliUC) DownloadDependency() {
	panic("ops")
}
