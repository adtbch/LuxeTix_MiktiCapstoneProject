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

// Helper function to handle error responses

// Get all submissions (pending events)
func (h *EventHandler) GetSubmissions(ctx echo.Context) error {
	var req dto.GetAllEventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	events, err := h.eventService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully All Pending Events", events))
}

// Get a single submission (pending event)
func (h *EventHandler) GetSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	event, err := h.eventService.GetByIDPending(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing a pending event", event))
}

// Get all events (general)
func (h *EventHandler) GetEvents(ctx echo.Context) error {
	var req dto.GetAllEventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	if req.MinPrice > 0 && req.MaxPrice > 0 && req.MinPrice > req.MaxPrice {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "minimum price cannot be greater than maximum price"))
	}

	events, err := h.eventService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing events", events))
}

// Get an event by ID
func (h *EventHandler) GetEvent(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	event, err := h.eventService.GetById(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing an events", event))
}

// Create an event


// Update event by user (with validation for user-specific updates)
func (h *EventHandler) UpdateEventByUser(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := h.eventService.UpdateEventbyUser(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully update an events", req))
}

// Update event by admin (perhaps allow admin to change other fields)
func (h *EventHandler) UpdateEventByAdmin(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	// Assuming different logic or validation for admins
	if err := h.eventService.UpdateEventbyUser(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully update an events", req))
}

// Delete an event
func (h *EventHandler) DeleteEvent(ctx echo.Context) error {
	var req dto.DeleteEventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	event, err := h.eventService.GetById(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if err := h.eventService.Delete(ctx.Request().Context(), event); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully deleted events", event))
}

func (h *EventHandler) GetAllEventByOwner(ctx echo.Context) error {
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	events, err := h.eventService.GetAllEventByOwner(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing events", events))
}
