package generate

import (
	"genos/internal/gen/base"
)

// GenerateBase - Генерация основного кода при инициализации макета
func GenerateBase(nameModule string) error {
	err := base.GenMain(nameModule)
	if err != nil {
		return err
	}
	err = base.GenApp()
	if err != nil {
		return err
	}
	err = base.GenOptionsHttpServer()
	if err != nil {
		return err
	}
	err = base.GenHttpServer()
	if err != nil {
		return err
	}
	return nil
}
