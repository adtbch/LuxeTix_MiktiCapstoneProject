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

type FilterMInMaxPrice struct {
	MinPrice int64 `param:"min_price" validate:"required"`
	MaxPrice int64 `param:"max_price" validate:"required"`
}

type FilterCategory struct {
	Category string `param:"category" validate:"required"`
}

type FilterLocation struct {
	Location string `param:"location" validate:"required"`	
}

type FilterPrice struct {
	Price int64 `param:"price" validate:"required"`
}

type FilterDate struct {
	Date string `param:"date" validate:"required"`
}

type FilterTime struct {
	Time string `param:"time" validate:"required"`
}

type SearchEvent struct {
	Keyword string `param:"keyword" validate:"required"`
}