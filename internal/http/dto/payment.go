package dto

type PaymentRequest struct {
	OrderID string `json:"order_id" validate:"required"`
	UserID  int64  `json:"user_id" validate:"required"`
	EventID int64  `json:"event_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Amount  int64  `json:"amount" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Email   string `json:"email" validate:"required"`
}
