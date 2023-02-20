package commands

type Generate struct {
}

// Do TODO
func (g *Generate) Do() error {
	panic("panic")
}

// GenerateRepository TODO
func (g *Generate) GenerateRepository() error {
	panic("panic")
}

// GenerateService TODO
func (g *Generate) GenerateService() error {
	panic("panic")
}

// GenerateController TODO
func (g *Generate) GenerateController() error {
	panic("panic")
}

func (g *Generate) GetNames() []string {
	return []string{"-g", "--generate"}
}
