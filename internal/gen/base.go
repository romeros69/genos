package generate

import (
	"genos/internal/gen/base"
)

// GenBase - Генерация основного кода при инициализации макета
func GenBase(nameModule string) error {
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
	err = base.GenPostgresOptions()
	if err != nil {
		return err
	}
	err = base.GenPostgres()
	if err != nil {
		return err
	}
	return nil
}
