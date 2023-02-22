package commands

import "fmt"

type Help struct {
}

// Do - TODO
func (h *Help) Do() error {
	fmt.Printf("Usage: genos [options] <arguments>\n\n" +
		"Options:\n" +
		"\t-h, --help\t\tList of options\n" +
		"\t-i, --init-layout\tGenerate project layout\n")
	return nil
}

func (h *Help) GetNames() []string {
	return []string{"-h", "--help"}
}
