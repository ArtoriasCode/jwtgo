package dto

type SecurityDTO struct {
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

type SignUpResponseDTO struct {
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Username string      `json:"username"`
	Gender   string      `json:"gender"`
	Security SecurityDTO `json:"security"`
}

type SignInRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
