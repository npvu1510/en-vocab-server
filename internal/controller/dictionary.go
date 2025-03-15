package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/service"
	"github.com/npvu1510/en-vocab-server/pkg/errors"
	"github.com/npvu1510/en-vocab-server/pkg/presenter/wrapper"
)

type IDictionaryController interface {
	GetDictionaries(ctx *fiber.Ctx, reqData dto.ListReqData) wrapper.Response
	// GetDictionariesWithCategoryId(ctx *fiber.Ctx, reqData dto.ListReqData) wrapper.Response
}

type DictionaryController struct {
	Service service.IDictionaryService
}

func NewDictionaryController(service service.IDictionaryService) IDictionaryController {
	return &DictionaryController{Service: service}
}

func (ctl *DictionaryController) GetDictionaries(ctx *fiber.Ctx, reqData dto.ListReqData) wrapper.Response {
	dictionaries, totalPages, err := ctl.Service.GetDictionaries(ctx.Context(), reqData)
	if err != nil {
		return wrapper.Response{
			Error: errors.ErrorInternalServer.Newf(err.Error()),
		}
	}

	// fmt.Printf("%v\n", dictionaries)

	return wrapper.Response{
		Error: errors.Success.New(),
		Data:  fiber.Map{"totalPages": totalPages, "dictionaries": dictionaries},
	}
}

// func (ctl *DictionaryController) GetDictionariesWithCategoryId(ctx *fiber.Ctx, reqData dto.ListReqData) wrapper.Response {
// 	dictionaries, totalPages, err := ctl.Service.GetDictionariesWithCategoryId(reqData)
// 	if err != nil {
// 		return wrapper.Response{
// 			Error: errors.ErrorInternalServer.Newf(err.Error()),
// 		}
// 	}

// 	// fmt.Printf("%v\n", dictionaries)

// 	return wrapper.Response{
// 		Error: errors.Success.New(),
// 		Data:  fiber.Map{"totalPages": totalPages, "dictionaries": dictionaries},
// 	}
// }
