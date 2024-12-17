package dto

type GetEventByIDRequest struct {
	ID int64 `param:"id" validate:"required"`
}
type DeleteEventRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type CreateEventRequest struct {
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Time          string `json:"time" validate:"required"`
	Date          string `json:"date" validate:"required"`
	Location      string `json:"location" validate:"required"`
	StatusEvent   string `json:"status event" validate:"required"`
	StatusRequest string `json:"status request" validate:"required"`
	Price         int64  `json:"price" validate:"required"`
	Category      string `json:"category" validate:"required"`
	UserID        int64  `json:"user_id" validate:"required"`
	Quantity      int64  `json:"quantity" validate:"required"`
}

type UpdateEventByAdminRequest struct {
	ID            int64  `param:"id" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Time          string `json:"time" validate:"required"`
	Location      string `json:"location" validate:"required"`
	StatusEvent   string `json:"status event" validate:"required"`
	StatusRequest string `json:"status request" validate:"required"`
	Price         int64  `json:"price" validate:"required"`
	Category      string `json:"category" validate:"required"`
	Date          string `json:"date" validate:"required"`
}

type UpdateEventByUserRequest struct {
	ID            int64  `param:"id" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Time          string `json:"time" validate:"required"`
	Location      string `json:"location" validate:"required"`
	StatusEvent   string `json:"status event" validate:"required"`
	StatusRequest string `json:"status request" validate:"required"`
	Price         int64  `json:"price" validate:"required"`
	Category      string `json:"category" validate:"required"`
	Date          string `json:"date" validate:"required"`
}
