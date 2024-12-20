package handler

import (
	"luxe/internal/http/dto"
	"luxe/internal/service"
	"luxe/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	tokenService service.TokenService
	userService service.UserService
}

func NewUserHandler(tokenService service.TokenService, userService service.UserService) UserHandler {
	return UserHandler{tokenService, userService}
}

func (h *UserHandler) Register(ctx echo.Context) error {
	var req dto.UserRegisterRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.userService.Register(ctx.Request().Context(), req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully create an account", nil))
}

func (h *UserHandler) Login(ctx echo.Context) error {
	var req dto.UserLoginRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	claims, err := h.userService.Login(ctx.Request().Context(), req.Username, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	token, err := h.tokenService.GennerateAccessToken(ctx.Request().Context(), *claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully login", map[string]string{"token": token}))
}