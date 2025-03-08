package cmd

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/npvu1510/en-vocab-server/internal"
	"github.com/npvu1510/en-vocab-server/internal/controller"
	"github.com/npvu1510/en-vocab-server/internal/repository"
	"github.com/npvu1510/en-vocab-server/internal/router"
	"github.com/npvu1510/en-vocab-server/internal/service"
	"github.com/npvu1510/en-vocab-server/pkg/config"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var StartServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		// internal.Invoke(startServerCmd).Run()
		internal.Invoke(startServerCmd).Start(context.Background())
	},
}

func startServerCmd(
	lc fx.Lifecycle,
	conf *config.Config,
	db *gorm.DB,
	dictionaryRouter *router.DictionaryRouter,
	dictionaryController controller.IDictionaryController,
	dictionaryService service.IDictionaryService,
	dictionaryRepo repository.IDictionaryRepo,

	categoryRouter *router.CategoryRouter,
	categoryController controller.ICategoryController,
	categoryService service.ICategoryService,
	categoryRepo repository.ICategoryRepo,
) {
	app := fiber.New()

	api := app.Group("/api")

	apiV1 := api.Group("/v1")
	{

		dictionaryRouter.RegisterRoutes(apiV1)
		categoryRouter.RegisterRoutes(apiV1)

	}

	app.Listen(fmt.Sprintf(":%v", conf.App.Port))

}
