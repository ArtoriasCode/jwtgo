package dto

type UserSignUpDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Gender   string `json:"gender" validate:"required,oneof=male female other"`
}

type UserSignInDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}
