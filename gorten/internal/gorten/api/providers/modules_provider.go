package providers

import (
	"gorten/internal/gorten/api/modules"

	"go.uber.org/fx"
)

func ModulesProvider() fx.Option {
	return fx.Options(
		modules.UserModule(),
		modules.GeneralModule(),
		modules.CategoryModule(),
		modules.CompanyModule(),
	)
}
