package dto

type SecurityDTO struct {
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Gender   string `json:"gender" validate:"required,oneof=male female other"`
}

type SignUpResponseDTO struct {
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Username string      `json:"username"`
	Gender   string      `json:"gender"`
	Security SecurityDTO `json:"security"`
}

type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}

type SignInResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignOutRequestDTO struct {
	Id string `json:"id"`
}

type SignOutResponseDTO struct {
	IsSignedOut bool `json:"is_signed_out"`
}

type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
