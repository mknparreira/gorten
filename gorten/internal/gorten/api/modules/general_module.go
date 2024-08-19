package modules

import (
	"gorten/internal/gorten/api/routes"

	"go.uber.org/fx"
)

func GeneralModule() fx.Option {
	return fx.Options(
		fx.Invoke(routes.GeneralRoute),
	)
}
