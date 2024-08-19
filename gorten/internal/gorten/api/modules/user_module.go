package modules

import (
	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/api/routes"
	"gorten/internal/gorten/repositories"
	"gorten/internal/gorten/services"

	"go.uber.org/fx"
)

func UserModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				func() string { return "marketplace" },
				fx.ResultTags(`name:"dbName"`),
			),
			fx.Annotate(
				func() string { return "users" },
				fx.ResultTags(`name:"collectionName"`),
			),

			fx.Annotate(
				repositories.UserRepositoryInit,
				fx.ParamTags(``, `name:"dbName"`, `name:"collectionName"`),
				fx.As(new(repositories.UserRepositoryImpl)),
			),
			fx.Annotate(
				services.UserServiceInit,
				fx.As(new(services.UserServiceImpl)),
			),
			handlers.User,
		),
		fx.Invoke(routes.UserRoute),
	)
}
