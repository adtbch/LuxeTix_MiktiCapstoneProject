package handler

// import (
// 	"net/http"

// 	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
// 	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
// 	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/response"
// 	validator "github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/validator"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/labstack/echo/v4"
// )

// type TransactionHandler struct {
// 	transactionService service.TransactionService
// 	paymentService     service.PaymentService
// }

// func NewTransactionHandler(transactionService service.TransactionService, paymentService service.PaymentService) *TransactionHandler {
// 	return &TransactionHandler{
// 		transactionService: transactionService,
// 		paymentService:     paymentService,
// 	}
// }

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
