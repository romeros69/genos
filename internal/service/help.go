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
		"\t-h, --help\t\t\t\tList of options\n" +
		"\t-i, --init-layout\t\t\tGenerate project layout\n" +
		"\t-g, --generate <path_to_dsl_file>\tGenerate CRUDL source code\n")
}
