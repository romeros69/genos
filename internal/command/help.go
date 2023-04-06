package command

import (
	"fmt"
	"genos/internal/service"
)

type Help struct {
	helpUC service.HelpContract
}

func NewHelp(helpUC service.HelpContract) *Help {
	return &Help{helpUC: helpUC}
}

func (h *Help) Do() error {
	fmt.Printf(h.helpUC.Help())
	return nil
}

func (h *Help) GetNames() []string {
	return []string{"-h", "--help"}
}
