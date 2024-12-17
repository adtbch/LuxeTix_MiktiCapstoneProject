package handler

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	service "github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/response"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventService service.EventService
	tokenService service.TokenService
}

func NewEventHandler(eventService service.EventService, tokenService service.TokenService) EventHandler {
	return EventHandler{eventService, tokenService}
}

func (h *EventHandler) GetSubmissions(ctx echo.Context) error {
	events, err := h.eventService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all pending events", events))
}

func (h *EventHandler) GetSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	event, err := h.eventService.GetByIDPending(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing an pending event", event))
}
func (h *EventHandler) GetEvents(ctx echo.Context) error {
	events, err := h.eventService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) GetEvent(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	event, err := h.eventService.GetById(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing an event", event))
}

func (h *EventHandler) CreateEvent(ctx echo.Context) error {
	// 1. Menggunakan helper function untuk mendapatkan UserID dari token
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, err.Error()))
	}

	// 2. Mengikat request body ke dalam dto
	var req dto.CreateEventRequest
	var tran dto.CreateEventTransactionRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// 3. Masukkan UserID dari token ke dalam request body
	req.UserID = userID

	// 4. Memanggil service untuk membuat event
	err = h.eventService.CreateEventByUser(ctx.Request().Context(), req, tran)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	// 5. Return success response
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully created an event", nil))
}

func (h *EventHandler) UpdateEventByUser(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.eventService.UpdateEventbyUser(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully updated an event", nil))
}

func (h *EventHandler) UpdateEventbyAdmin(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.eventService.UpdateEventbyUser(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully updated an event", nil))
}

func (h *EventHandler) DeleteEvent(ctx echo.Context) error {
	var req dto.DeleteEventRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	event, err := h.eventService.GetById(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.eventService.Delete(ctx.Request().Context(), event)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully deleted an event", nil))
}
