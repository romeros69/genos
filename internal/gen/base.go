package generate

import (
	"fmt"
	"genos/internal/gen/base"
)

// GenBase - Генерация основного кода при инициализации макета
func GenBase(moduleName string) error {
	generators := initGenerators(moduleName)

	for _, v := range generators {
		err := v.GenerateCode()
		if err != nil {
			return fmt.Errorf("error in GenerateCode(): %w", err)
		}
	}
	return nil
}

func initGenerators(moduleName string) []base.Generator {
	return []base.Generator{
		0: base.NewMainGenerator(moduleName),
		1: base.NewAppGenerator(moduleName),
		2: base.NewConfigGenerator(moduleName),
		3: base.NewPostgresGenerator(moduleName),
		4: base.NewPostgresOptionGenerator(moduleName),
		5: base.NewHttpServerGenerator(moduleName),
		6: base.NewHttpOptionsGenerator(moduleName),
	}
}
