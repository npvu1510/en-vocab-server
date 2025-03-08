package internal

import (
	"github.com/npvu1510/en-vocab-server/internal/controller"
	"github.com/npvu1510/en-vocab-server/internal/repository"
	"github.com/npvu1510/en-vocab-server/internal/router"
	"github.com/npvu1510/en-vocab-server/internal/service"
	"github.com/npvu1510/en-vocab-server/pkg/config"
	"go.uber.org/fx"
)

func Invoke(invokers ...any) *fx.App {
	var conf = config.MustLoad()

	return fx.New(
		fx.Provide(
			// db
			NewDatabaseConnection,

			// routers
			router.NewDictionaryRouter,
			router.NewCategoryRouter,

			// controllers
			controller.NewDictionaryController,
			controller.NewCategoryController,

			// services
			service.NewDictionaryService,
			service.NewCategoryService,

			// repos
			repository.NewDictionaryRepo,
			repository.NewCategoryRepo,
		),
		fx.Supply(conf),
		fx.Invoke(invokers...))
}
