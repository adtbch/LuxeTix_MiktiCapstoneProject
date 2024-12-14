package entity

type Transaction struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	EventID int64  `json:"event_id"`
	Status  string `json:"status"`
	CreatedAt int  `json:"created_at"`
	UpdatedAt int  `json:"updated_at"`
}

