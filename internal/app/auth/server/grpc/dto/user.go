package dto

type UserDTO struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
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
}
