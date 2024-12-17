package entity

type Payment struct {
	OrderID string
	UserID  int64
	Amount  int64
	Status  string
	Email   string
}
