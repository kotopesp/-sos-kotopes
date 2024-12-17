package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/seeker"
	"github.com/kotopesp/sos-kotopes/internal/core"
	"github.com/kotopesp/sos-kotopes/pkg/logger"
)

//	@Summary		Get seeker
//	@Tags			seeker
//	@Description	Get seeker by id
//	@ID				get-seeker
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	model.Response{data=seeker.ResponseSeeker}
//	@Failure		400		{object}	model.Response
//	@Failure		500		{object}	model.Response
//	@Security		ApiKeyAuthBasic
//	@Router			/seekers/{user_id}  [get]
func (r *Router) getSeeker(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("user_id")
	if err != nil {
		logger.Log().Debug(ctx.Context(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}

	getSeeker, err := r.seekerService.GetSeeker(ctx.UserContext(), id)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrSeekerNotFound):
			logger.Log().Debug(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse(err.Error()))
		default:
			logger.Log().Error(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
		}
	}

	responseSeeker := seeker.ToResponseSeeker(&getSeeker)
	return ctx.Status(fiber.StatusOK).JSON(responseSeeker)
}

//	@Summary		Create a seeker
//	@Tags			seeker
//	@Description	Create a seeker
//	@ID				create-seeker
//	@Accept			json
//	@Produce		json
//	@Param			request	body		seeker.CreateSeeker	true	"Seeker"
//	@Success		200		{object}	model.Response{data=seeker.ResponseSeeker}
//	@Failure		400		{object}	model.Response
//	@Failure		500		{object}	model.Response
//	@Security		ApiKeyAuthBasic
//	@Router			/seekers [post]
func (r *Router) createSeeker(ctx *fiber.Ctx) error {
	var createSeeker seeker.CreateSeeker

	fiberError, parseOrValidationError := parseBodyAndValidate(ctx, r.formValidator, &createSeeker)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}

	equipment, err := createSeeker.GetEquipment()
	if err != nil {
		logger.Log().Debug(ctx.Context(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}

	coreSeeker, err := r.seekerService.CreateSeeker(ctx.UserContext(), createSeeker.ToCoreSeeker(), equipment)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrNoSuchUser):
			logger.Log().Debug(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse(err.Error()))
		default:
			logger.Log().Error(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
		}
	}

	responseSeeker := seeker.ToResponseSeeker(&coreSeeker)
	return ctx.Status(fiber.StatusOK).JSON(responseSeeker)
}

//	@Summary		Update a seeker
//	@Tags			seeker
//	@Description	Update a seeker
//	@ID				update-seeker
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		seeker.UpdateSeeker	true	"Seeker"
//
//	@Success		200		{object}	model.Response{data=seeker.ResponseSeeker}
//	@Failure		400		{object}	model.Response
//	@Failure		500		{object}	model.Response
//	@Security		ApiKeyAuthBasic
//	@Router			/seekers/{user_id}  [patch]
func (r *Router) updateSeeker(ctx *fiber.Ctx) error {
	//userID, err := getIDFromToken(ctx)
	//if err != nil {
	//	logger.Log().Debug(ctx.UserContext(), err.Error())
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
	//}

	var update seeker.UpdateSeeker

	fiberError, parseOrValidationError := parseBodyAndValidate(ctx, r.formValidator, &update)
	if fiberError != nil || parseOrValidationError != nil {
		return fiberError
	}
	coreUpdate := update.ToCoreUpdateSeeker()

	updatedSeeker, err := r.seekerService.UpdateSeeker(ctx.UserContext(), coreUpdate)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrNoSuchUser):
			logger.Log().Debug(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse(err.Error()))
		case errors.Is(err, core.ErrEmptyUpdateRequest):
			logger.Log().Debug(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
		default:
			logger.Log().Error(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
		}
	}
	responseUser := seeker.ToResponseSeeker(&updatedSeeker)

	return ctx.Status(fiber.StatusOK).JSON(responseUser)
}

//	@Summary		Delete seeker
//	@Tags			seeker
//	@Description	Delete seeker
//	@ID				delete-seeker
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	model.Response{data=seeker.ResponseSeeker}
//	@Failure		400		{object}	model.Response
//	@Failure		500		{object}	model.Response
//	@Security		ApiKeyAuthBasic
//	@Router			/seekers/{user_id}  [delete]
func (r *Router) deleteSeeker(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("user_id")
	if err != nil {
		logger.Log().Debug(ctx.Context(), err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse(err.Error()))
	}

	if err = r.seekerService.DeleteSeeker(ctx.UserContext(), id); err != nil {
		switch {
		case errors.Is(err, core.ErrSeekerNotFound):
			logger.Log().Debug(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse(err.Error()))
		default:
			logger.Log().Error(ctx.UserContext(), err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse(err.Error()))
		}
	}

	return ctx.Status(fiber.StatusOK).JSON("Delete")
}
