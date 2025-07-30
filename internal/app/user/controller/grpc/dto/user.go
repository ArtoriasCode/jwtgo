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

type UserIdDTO struct {
	Id string `json:"id"`
}

type UserEmailDTO struct {
	Email string `json:"email" validate:"required,email"`
}

type UserCreateDTO struct {
	Email    string      `json:"email" validate:"required,email"`
	Role     string      `json:"role"`
	Security SecurityDTO `json:"security"`
}

type UserUpdateDTO struct {
	Id       string      `json:"id"`
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Security SecurityDTO `json:"security"`
}
