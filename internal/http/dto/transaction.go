package dto

type CreateEventTransactionRequest struct {
	UserID  int64   `json:"user_id" validate:"required"`
}