package modules

import (
	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/api/routes"
	"gorten/internal/gorten/repositories"
	"gorten/internal/gorten/services"

	"go.uber.org/fx"
)

func ProductModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				func() string { return "products" },
				fx.ResultTags(`name:"collectionProductName"`),
			),
			fx.Annotate(
				repositories.ProductRepositoryInit,
				fx.ParamTags(``, ``, `name:"collectionProductName"`),
				fx.As(new(repositories.ProductRepositoryImpl)),
			),
			fx.Annotate(
				services.ProductServiceInit,
				fx.As(new(services.ProductServiceImpl)),
			),
			handlers.Product,
		),
		fx.Invoke(routes.ProductRoute),
	)
}
