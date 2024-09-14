package modules

import (
	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/api/routes"
	"gorten/internal/gorten/repositories"
	"gorten/internal/gorten/services"

	"go.uber.org/fx"
)

func CategoryModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				func() string { return "categories" },
				fx.ResultTags(`name:"collectionCategoryName"`),
			),
			fx.Annotate(
				repositories.CategoryRepositoryInit,
				fx.ParamTags(``, ``, `name:"collectionCategoryName"`),
				fx.As(new(repositories.CategoryRepositoryImpl)),
			),
			fx.Annotate(
				services.CategoryServiceInit,
				fx.As(new(services.CategoryServiceImpl)),
			),
			handlers.Category,
		),
		fx.Invoke(routes.CategoryRoute),
	)
}
