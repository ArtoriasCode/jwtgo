package dto

type SignUpRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Gender   string `json:"gender" validate:"required,oneof=male female other"`
}

type SignUpResponseDTO struct {
	Message string `json:"message"`
}

type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}

type SignInResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
}

type SignOutRequestDTO struct {
	Id string `json:"id" validate:"required"`
}

type SignOutResponseDTO struct {
	Message string `json:"message"`
}

type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
}
