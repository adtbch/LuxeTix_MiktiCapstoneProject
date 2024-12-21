package dto

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetUserByIDRequest struct {
	ID int64 `param:"ID" validate:"required"`
}

type CreateUserByRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Role string `json:"role" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UpdateUserRequest struct {
	ID int64 `param:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Role string `json:"role" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type ResetPasswordRequest struct {
	Token string `param:"reset_password_token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyEmailRequest struct {
	Token string `param:"verify_email_token" validate:"required"`
}

type RequestResetPassword struct {
	Username string `json:"username" validate:"required"`
}