package command

import (
	"fmt"
	"genos/internal/service"
)

type Help struct {
	helpUC service.HelpUC
}

func (h *Help) Do() error {
	fmt.Printf(h.helpUC.Help())
	return nil
}

func (h *Help) GetNames() []string {
	return []string{"-h", "--help"}
}
