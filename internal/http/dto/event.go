package dto

type GetEventByIDRequest struct {
	ID int64 `param:"id" validate:"required"`
	UserID int64 `json:"user_id" validate:"required"`
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
	Quantity      int64  `json:"quantity" validate:"required"`
}

type UpdateEventByUserRequest struct {
	ID            int64  `param:"id" validate:"required"`
	UserID        int64  `json:"user_id" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Time          string `json:"time" validate:"required"`
	Location      string `json:"location" validate:"required"`
	StatusEvent   string `json:"status event" validate:"required"`
	StatusRequest string `json:"status request" validate:"required"`
	Price         int64  `json:"price" validate:"required"`
	Category      string `json:"category" validate:"required"`
	Date          string `json:"date" validate:"required"`
	Quantity      int64  `json:"quantity" validate:"required"`
}

type GetAllEventRequest struct {
	Search string `query:"search"`
	Filter string `query:"filter"`
	Sort   string `query:"sort"`
	Order  string `query:"order"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`

	// Tambahan filter dari kode pertama
	MinPrice      float64 `query:"min_price"`
	MaxPrice      float64 `query:"max_price"`
	Category      string  `query:"category"`
	Location      string  `query:"location"`
	Time          string  `query:"time"`
	Date          string  `query:"date"`
	StatusEvent   string  `query:"status_event"`
	StatusRequest string  `query:"status_request"`
	Price         float64 `query:"price"`
}
