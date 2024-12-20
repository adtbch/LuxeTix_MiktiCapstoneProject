package dto

type CreateEventTransactionRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
}

type CreateOrderRequest struct {
	OrderID  int64  `json:"id" validate:"required"`
	UserID   int64  `json:"user_id" validate:"required"`
	EventID  int64  `json:"event_id" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
	Amount   int64  `json:"amount" validate:"required"`
	Status   string `json:"status" validate:"required"`
}