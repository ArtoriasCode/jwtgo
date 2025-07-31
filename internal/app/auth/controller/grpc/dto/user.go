package dto

type UserTokenDTO struct {
	Token string `json:"token"`
}

type UserTokensDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Gender   string `json:"gender" validate:"required,oneof=male female other"`
}

type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}
