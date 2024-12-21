package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/dto"
	service "github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/response"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionService
	tokenService       service.TokenService
	paymentService     service.PaymentService
	userService        service.UserService
	eventService       service.EventService
}

func NewTransactionHandler(transactionService service.TransactionService, tokenService service.TokenService, paymentService service.PaymentService, userService service.UserService, eventService service.EventService) TransactionHandler {
	return TransactionHandler{transactionService, tokenService, paymentService, userService, eventService}
}

func (h *TransactionHandler) CreateOrder(ctx echo.Context) error {
    // Extract user ID from token
    userID, err := h.tokenService.ExtractUserIDFromToken(ctx)
    if err != nil {
        fmt.Println("Failed to extract user ID:", err)
        return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
    }

    // Retrieve user data
    user, err := h.userService.GetById(ctx.Request().Context(), userID)
    if err != nil {
        fmt.Println("Failed to retrieve user:", err)
        return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
    }

    // Bind request body to CreateOrderRequest
    var req dto.CreateOrderRequest
    if err := ctx.Bind(&req); err != nil {
        fmt.Println("Failed to bind request:", err)
        return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
    }
    req.UserID = userID

    // Retrieve event details using EventID from the request
    event, err := h.eventService.GetById(ctx.Request().Context(), req.EventID)
    if err != nil {
        fmt.Println("Failed to retrieve event:", err)
        return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
    }

    err, transaction := h.transactionService.CreateOrder(ctx.Request().Context(), req)
    if err != nil {
        fmt.Println("Failed to create order:", err)
        return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
    }

    req.Amount = event.Price * req.Quantity
    req.OrderID = transaction.ID

    var paymentRequest dto.PaymentRequest
    // Prepare paymentRequest details
    paymentRequest.OrderID = strconv.FormatInt(req.OrderID, 10) // Assuming req.OrderID is int64
    paymentRequest.UserID = userID
    paymentRequest.EventID = req.EventID
    paymentRequest.Title = event.Title
    paymentRequest.Amount = req.Amount
    paymentRequest.Status = req.Status
    paymentRequest.Email = user.Email

    // Create payment transaction
    payment, err := h.paymentService.CreateTransaction(ctx.Request().Context(), paymentRequest, user)
    if err != nil {
        fmt.Println("Failed to create payment transaction:", err)
        return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
    }

    // Return success response with created order information
    return ctx.JSON(http.StatusCreated, response.SuccessResponse("Successfully created an order", payment))
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

func (h *TransactionHandler) WebHookTransaction(ctx echo.Context) error {
	var input entity.MidtransRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	orderID, err := strconv.ParseInt(input.OrderID, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	transaction, err := h.transactionService.GetById(ctx.Request().Context(), orderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	if input.TransactionStatus == "settlement" {
		status := "paid"
		err = h.transactionService.Update(ctx.Request().Context(), orderID, status)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		event, err := h.transactionService.GetById(ctx.Request().Context(), orderID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		if event.Type == "tiket"{

		}
		transaction.Status = status
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("Successfully Payment", nil))
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
