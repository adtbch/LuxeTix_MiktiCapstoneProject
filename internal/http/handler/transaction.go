package handler

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	service "github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/response"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionService
	tokenService       service.TokenService
}

func NewTransactionHandler(transactionService service.TransactionService, tokenService service.TokenService) TransactionHandler {
	return TransactionHandler{transactionService, tokenService}
}

func (h *TransactionHandler) CreateOrder(ctx echo.Context) error {
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	var req dto.CreateOrderRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	req.UserID = userID

	err = h.transactionService.CreateOrder(ctx.Request().Context(), req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse("Successfully created an order", nil))
}

func (h *TransactionHandler) GetUserTransactions(ctx echo.Context) error {
	userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	transactions, err := h.transactionService.GetUserTransactions(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully get user transactions", transactions))

}

// func (h *TransactionHandler) CreateOrder(ctx echo.Context) error {
// 	var input struct {
// 		OrderID string `json:"order_id" validate:"required"`
// 		Amount  int64  `json:"amount" validate:"required"`
// 	}

// 	if err := ctx.Bind(&input); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
// 	}

// 	dataUser, _ := ctx.Get("user").(*jwt.Token)
// 	claims := dataUser.Claims.(*common.JwtCustomClaims)

// 	transaction := entity.Transaction(input.OrderID, claims.ID, input.Amount, "unpaid")

// 	err := h.transactionService.Create(ctx.Request().Context(), transaction)

// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	paymentRequest := entity.NewPaymentRequest(transaction.OrderID, transaction.Amount, claims.Name, "", claims.Email)

// 	payment, err := h.paymentService.CreateOrder(ctx.Request().Context(), paymentRequest)

// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	return ctx.JSON(http.StatusCreated, map[string]string{"url_pembayaran": payment})
// }
