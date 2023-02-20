package commands

type Generate struct {
}

func (g *Generate) GetNames() []string {
	return []string{"-g", "--generate"}
}
