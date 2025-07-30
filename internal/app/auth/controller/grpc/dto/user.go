package dto

type UserTokenDTO struct {
	Token string `json:"token"`
}

type UserTokensDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refresh_token"`
}

type UserSignUpDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type UserSignInDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}
