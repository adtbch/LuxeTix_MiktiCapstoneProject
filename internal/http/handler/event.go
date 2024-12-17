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
func (h *EventHandler) handleErrorResponse(ctx echo.Context, status int, err error) error {
	return ctx.JSON(status, response.ErrorResponse(status, err.Error()))
}

// Helper function to handle success responses
func (h *EventHandler) handleSuccessResponse(ctx echo.Context, status int, message string, data interface{}) error {
	return ctx.JSON(status, response.SuccessResponse(message, data))
}

// Get all submissions (pending events)
func (h *EventHandler) GetSubmissions(ctx echo.Context) error {
	events, err := h.eventService.GetAll(ctx.Request().Context())
	if err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}
	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully showing all pending events", events)
}

// Get a single submission (pending event)
func (h *EventHandler) GetSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusBadRequest, err)
	}

	event, err := h.eventService.GetByIDPending(ctx.Request().Context(), req.ID)
	if err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}
	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully showing a pending event", event)
}

// Get all events (general)
func (h *EventHandler) GetEvents(ctx echo.Context) error {
	events, err := h.eventService.GetAll(ctx.Request().Context())
	if err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}
	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully showing all events", events)
}

// Get an event by ID
func (h *EventHandler) GetEvent(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusBadRequest, err)
	}

	event, err := h.eventService.GetById(ctx.Request().Context(), req.ID)
	if err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully showing an event", event)
}

// Create an event
func (h *EventHandler) CreateEvent(ctx echo.Context) error {
	// Extract UserID from token
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return h.handleErrorResponse(ctx, http.StatusUnauthorized, err)
	}

	var req dto.CreateEventRequest
	if err := ctx.Bind(&req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusBadRequest, err)
	}

	// Assign UserID to the request
	req.UserID = userID

	// Create the event
	if err := h.eventService.CreateEventByUser(ctx.Request().Context(), req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully created an event", nil)
}

// Update event by user (with validation for user-specific updates)
func (h *EventHandler) UpdateEventByUser(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest
	if err := ctx.Bind(&req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusBadRequest, err)
	}

	if err := h.eventService.UpdateEventbyUser(ctx.Request().Context(), req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully updated an event", nil)
}

// Update event by admin (perhaps allow admin to change other fields)
func (h *EventHandler) UpdateEventByAdmin(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest
	if err := ctx.Bind(&req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusBadRequest, err)
	}

	// Assuming different logic or validation for admins
	if err := h.eventService.UpdateEventbyUser(ctx.Request().Context(), req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully updated event by admin", nil)
}

// Delete an event
func (h *EventHandler) DeleteEvent(ctx echo.Context) error {
	var req dto.DeleteEventRequest
	if err := ctx.Bind(&req); err != nil {
		return h.handleErrorResponse(ctx, http.StatusBadRequest, err)
	}

	event, err := h.eventService.GetById(ctx.Request().Context(), req.ID)
	if err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	if err := h.eventService.Delete(ctx.Request().Context(), event); err != nil {
		return h.handleErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return h.handleSuccessResponse(ctx, http.StatusOK, "Successfully deleted an event", nil)
}


func (h *EventHandler) SortFromCheapestToExpensivest(ctx echo.Context) error {
	events, err := h.eventService.SortFromCheapestToExpensivest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) SortFromExpensivestToCheapest(ctx echo.Context) error {
	events, err := h.eventService.SortFromExpensivestToCheapest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) SortNewestToOldest(ctx echo.Context) error {
	events, err := h.eventService.SortNewestToOldest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) FilteringByCategory(ctx echo.Context) error {
	var req dto.FilterCategory

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	events, err := h.eventService.FilteringByCategory(ctx.Request().Context(), req.Category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) FilteringByLocation(ctx echo.Context) error {
	var req dto.FilterLocation

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	events, err := h.eventService.FilteringByLocation(ctx.Request().Context(), req.Location)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) FilteringByDate(ctx echo.Context) error {
	var req dto.FilterDate

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	events, err := h.eventService.FilteringByDate(ctx.Request().Context(), req.Date)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) FilteringByPrice(ctx echo.Context) error {
	var req dto.FilterPrice

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	events, err := h.eventService.FilteringByPrice(ctx.Request().Context(), req.Price)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}

func (h *EventHandler) FilteringMInMaxPrice(ctx echo.Context) error {
	var req dto.FilterMInMaxPrice

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	events, err := h.eventService.FilterMaxMinPrice(ctx.Request().Context(), req.MinPrice, req.MaxPrice)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all events", events))
}
