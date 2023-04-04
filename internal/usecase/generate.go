package usecase

type GenerateSource struct {
	fw FileSourceWorker
}

func NewGenerateSource(fw FileSourceWorker) *GenerateSource {
	return &GenerateSource{fw: fw}
}

// GenerateBaseCode Генерация базового кода
func (gs *GenerateSource) GenerateBaseCode() {

}

// GenerateCRUDL Генерация основного кода из DSL
func (gs *GenerateSource) GenerateCRUDL() {
	panic("mock")
}
