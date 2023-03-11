package generate

import (
	"genos/internal/gen/base"
)

// GenBase - Генерация основного кода при инициализации макета
func GenBase(moduleName string) error {
	err := base.GenMain(moduleName)
	if err != nil {
		return err
	}
	err = base.GenApp(moduleName)
	if err != nil {
		return err
	}
	err = base.GenConfig()
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
	err = base.GenPostgres(moduleName)
	if err != nil {
		return err
	}
	return nil
}
