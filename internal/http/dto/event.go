package dto

type GetEventByIDRequest struct {
	ID int64 `param:"id" validate:"required"`
}
type DeleteEventRequest struct {
	ID int64 `param:"id" validate:"required"`
}

type CreateEventRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Time        string  `json:"time" validate:"required"`
	Date        string  `json:"date" validate:"required"`
	Location    string  `json:"location" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	UserID      int     `json:"user_id" validate:"required"`
}

type UpdateEventRequest struct {
	ID          int64   `param:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Time        string  `json:"time" validate:"required"`
	Location    string  `json:"location" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	Date        string  `json:"date" validate:"required"`
}
