package entity

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	ResetPasswordToken string `json:"reset_password_token"`
	VerifyEmailToken string `json:"verify_email_token"`
	IsVerified bool `json:"is_verified"`
}

func (User) TableName() string {
	return "users"
}