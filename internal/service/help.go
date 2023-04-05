package service

import "fmt"

type HelpUC struct {
}

func NewHelpUseCase() *HelpUC {
	return &HelpUC{}
}

var _ HelpContract = (*HelpUC)(nil)

func (h HelpUC) Help() string {
	return fmt.Sprintf("Usage: genos [options] <arguments>\n\n" +
		"Options:\n" +
		"\t-h, --help\t\tList of options\n" +
		"\t-i, --init-layout\tGenerate project layout\n")
}
