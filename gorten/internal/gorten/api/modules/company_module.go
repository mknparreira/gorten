package modules

import (
	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/api/routes"
	"gorten/internal/gorten/repositories"
	"gorten/internal/gorten/services"

	"go.uber.org/fx"
)

func CompanyModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				func() string { return "companies" },
				fx.ResultTags(`name:"collectionCompanyName"`),
			),
			fx.Annotate(
				repositories.CompanyRepositoryInit,
				fx.ParamTags(``, ``, `name:"collectionCompanyName"`),
				fx.As(new(repositories.CompanyRepositoryImpl)),
			),
			fx.Annotate(
				services.CompanyServiceInit,
				fx.As(new(services.CompanyServiceImpl)),
			),
			handlers.Company,
		),
		fx.Invoke(routes.CompanyRoute),
	)
}
