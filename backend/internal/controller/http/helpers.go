package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/validator"
	"github.com/kotopesp/sos-kotopes/pkg/logger"
)

// token helpers: getting info from token
func getIDFromToken(ctx *fiber.Ctx) (id int, err error) {
	idItem := getPayloadItem(ctx, "id")

	logger.Log().Debug(ctx.UserContext(), fmt.Sprintf("idItem: %v", idItem))

	idFloat, ok := idItem.(float64)
	if !ok {
		return 0, model.ErrInvalidTokenID
	}
	return int(idFloat), nil
}

func getUsernameFromToken(ctx *fiber.Ctx) (username string, err error) {
	usernameItem := getPayloadItem(ctx, "username")
	username, ok := usernameItem.(string)
	if !ok {
		return "", model.ErrInvalidTokenUsername
	}
	return username, nil
}

// validation helpers
func parseBodyAndValidate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	if err := ctx.BodyParser(data); err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error())), err
	}

	return validate(ctx, formValidator, data)
}

func parseQueryAndValidate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	if err := ctx.QueryParser(data); err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error())), err
	}

	return validate(ctx, formValidator, data)
}

func parseParamsAndValidate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	if err := ctx.ParamsParser(data); err != nil {
		logger.Log().Error(ctx.UserContext(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error())), err
	}

	return validate(ctx, formValidator, data)
}

// helper for parse...AndValidate
func validate(ctx *fiber.Ctx, formValidator validator.FormValidatorService, data interface{}) (fiberError, parseOrValidationError error) {
	errs := formValidator.Validate(data)
	if len(errs) > 0 {
		logger.Log().Info(ctx.UserContext(), fmt.Sprintf("%v", errs))
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.ErrorResponse(fiber.Map{
			"validation_errors": errs,
		})), model.ErrValidationFailed
	}
	return nil, nil
}
