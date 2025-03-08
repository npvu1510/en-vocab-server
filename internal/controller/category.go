package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/service"
	"github.com/npvu1510/en-vocab-server/pkg/errors"
	"github.com/npvu1510/en-vocab-server/pkg/presenter/wrapper"
)

type ICategoryController interface {
	GetCategories(ctx *fiber.Ctx, reqData dto.ListReqData) wrapper.Response
}

type CategoryController struct {
	Service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) ICategoryController {
	return &CategoryController{Service: service}
}

func (ctl *CategoryController) GetCategories(ctx *fiber.Ctx, reqData dto.ListReqData) wrapper.Response {
	categories, err := ctl.Service.GetCategories(reqData)
	if err != nil {
		return wrapper.Response{
			Error: errors.ErrorInternalServer.Newf(err.Error()),
		}
	}

	return wrapper.Response{
		Error: errors.Success.New(),
		Data:  categories,
	}
}
