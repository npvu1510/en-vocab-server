package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/npvu1510/en-vocab-server/internal/controller"
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/pkg/presenter/wrapper"
)

type CategoryRouter struct {
	Ctl controller.ICategoryController
}

func NewCategoryRouter(ctl controller.ICategoryController) *CategoryRouter {
	return &CategoryRouter{Ctl: ctl}
}

func (r *CategoryRouter) RegisterRoutes(apiRouter fiber.Router) {
	categoriesRoute := apiRouter.Group("/categories")
	categoriesRoute.Get("/", wrapper.Wrapper[dto.ListReqData](r.Ctl.GetCategories))

}
