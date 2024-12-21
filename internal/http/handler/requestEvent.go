package handler

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	service "github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/response"

	"github.com/labstack/echo/v4"
)

type RequestEventHandler struct {
	RequestEventService service.RequestEventService
	tokenService      service.TokenService
}

func NewRequestEventHandler(requestEventService service.RequestEventService, tokenService service.TokenService) RequestEventHandler {
	return RequestEventHandler{requestEventService, tokenService}

}

func (h *RequestEventHandler) GetSubmissions(ctx echo.Context) error {
	var req dto.GetAllEventsSubmission
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	submission, err := h.RequestEventService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing all submission", submission))
}

func (h *RequestEventHandler) GetSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	submission, err := h.RequestEventService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully showing a submission", submission))
}

func (h *RequestEventHandler) UpdateSubmissionByUser(ctx echo.Context) error {
	var req dto.UpdateEventByUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, err.Error()))
	}
	req.UserID = userID
	err = h.RequestEventService.UpdateByUser(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully update a submission", nil))
}

func (h *RequestEventHandler) ApproveSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	err := h.RequestEventService.Approve(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully approve a submission", nil))
}

func (h *RequestEventHandler) RejectSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	err := h.RequestEventService.Reject(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully reject a submission", nil))
}

func (h *RequestEventHandler) CancelSubmission(ctx echo.Context) error {
	var req dto.GetEventByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, err.Error()))
	}
	req.UserID = userID
	submission, err := h.RequestEventService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	err = h.RequestEventService.Cancel(ctx.Request().Context(), submission, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully cancel a submission", nil))
}

func (h *RequestEventHandler) CreateEvent(ctx echo.Context) error {
	// Extract UserID from token
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	var req dto.CreateEventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	// Assign UserID to the request
	req.UserID = userID

	// Create the event
	if err := h.RequestEventService.Create(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully created an events", req))
}