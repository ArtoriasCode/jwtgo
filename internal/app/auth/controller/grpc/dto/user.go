package dto

type SecurityDTO struct {
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	RefreshToken string `json:"refresh_token"`
}

type UserDTO struct {
	Id        string      `json:"id"`
	Email     string      `json:"email"`
	Role      string      `json:"role"`
	Security  SecurityDTO `json:"security"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
}

type UserTokenDTO struct {
	Token string `json:"token"`
}

type UserTokensDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refresh_token"`
}

type UserCredentialsDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}
