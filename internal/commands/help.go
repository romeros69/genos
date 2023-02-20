package commands

type Help struct {
}

func (h *Help) GetNames() []string {
	return []string{"-h", "--help"}
}
