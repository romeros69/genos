package commands

type Help struct {
}

// Do - TODO
func (h *Help) Do() error {
	panic("panic")
}

func (h *Help) GetNames() []string {
	return []string{"-h", "--help"}
}
