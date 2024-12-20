package handler

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	service "github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/response"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationService service.NotificationService
	tokenService        service.TokenService
}

func NewNotificationHandler(notificationService service.NotificationService, tokenService service.TokenService) NotificationHandler {
	return NotificationHandler{notificationService, tokenService}
}

// GetAllNotification
func (h *NotificationHandler) GetAllNotification(c echo.Context) error {
	Notifications, err := h.notificationService.GetAllNotification(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Notifications,
	})
}

// func untuk create notification
func (h *NotificationHandler) CreateNotification(ctx echo.Context) error {
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	var input dto.NotificationInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	input.UserID = userID

	if err := h.notificationService.CreateNotification(ctx.Request().Context(), input); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully created an events", input))
}

// get notification after get chage value isRead to true and only get notification if isread false
func (h *NotificationHandler) UserGetNotification(ctx echo.Context) error {
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	var input dto.NotificationInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	input.UserID = userID

	Notifications, err := h.notificationService.UserGetNotification(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": Notifications,
	})
}
