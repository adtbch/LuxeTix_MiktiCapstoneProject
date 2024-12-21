package dto

type GetAllEventsSubmission struct {
	Page   int  `query:"page" validate:"required"`
	Limit  int  `query:"limit" validate:"required"`
	Search string `query:"search" validate:"required"`
	Sort   string `query:"sort" validate:"required"`
	Order  string `query:"order" validate:"required"`
}