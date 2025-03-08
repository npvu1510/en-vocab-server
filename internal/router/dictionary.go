package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/npvu1510/en-vocab-server/internal/controller"
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/pkg/presenter/wrapper"
)

type DictionaryRouter struct {
	Ctl controller.IDictionaryController
}

func NewDictionaryRouter(ctl controller.IDictionaryController) *DictionaryRouter {
	return &DictionaryRouter{Ctl: ctl}
}

func (r *DictionaryRouter) RegisterRoutes(apiRouter fiber.Router) {
	dictionariesRoute := apiRouter.Group("/dictionaries")

	dictionariesRoute.Get("/", wrapper.Wrapper[dto.ListReqData](r.Ctl.GetDictionaries))
	dictionariesRoute.Post("/", wrapper.Wrapper[dto.ListReqData](r.Ctl.GetDictionaries))
	dictionariesRoute.Put("/", wrapper.Wrapper[dto.ListReqData](r.Ctl.GetDictionaries))

}
