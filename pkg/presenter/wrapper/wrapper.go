package wrapper

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/npvu1510/en-vocab-server/pkg/errors"
)

const (
	DataField    = "data"
	TraceIDField = "trace_id"
	StatusField  = "status"
	CodeField    = "code"
	MessageField = "message"
)

type Response struct {
	Error error
	Data  any
}

type handlerWithReqData[RequestDataType any] func(ctx *fiber.Ctx, reqData RequestDataType) Response

func Wrapper[RequestDataType any](handler handlerWithReqData[RequestDataType]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var reqData RequestDataType

		if err := ctx.QueryParser(&reqData); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		if ctx.Method() == "POST" || ctx.Method() == "PUT" || ctx.Method() == "PATCH" {
			if err := ctx.BodyParser(&reqData); err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
			}
		}

		if err := ctx.ParamsParser(&reqData); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}

		res := handler(ctx, reqData)
		// fmt.Printf("RESPONSE ERROR: %+v\n", res.Error)
		// fmt.Printf("RESPONSE DATA: %+v\n", res.Data)

		return FormatAPIResponse(ctx, res)

	}
}

// FormatAPIResponse presenter.Response to response
func FormatAPIResponse(c *fiber.Ctx, res Response) error {
	result := fiber.Map{}
	if _, ok := res.Error.(errors.CustomError); ok {
		status := int(errors.GetType(res.Error))
		result[StatusField] = status
		result[MessageField] = res.Error.Error() // errors.GetMsg(status)
		result[CodeField] = errors.GetCode(status)
	}

	if res.Data != nil {
		result[DataField] = res.Data
	}

	return c.Status(http.StatusOK).JSON(result)
}
