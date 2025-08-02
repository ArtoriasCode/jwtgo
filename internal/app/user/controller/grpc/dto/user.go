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
	Username  string      `json:"username"`
	Gender    string      `json:"gender"`
	Security  SecurityDTO `json:"security"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
}

type GetByIdRequestDTO struct {
	Id string `json:"id"`
}

type GetByEmailRequestDTO struct {
	Email string `json:"email"`
}

type CreateRequestDTO struct {
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Username string      `json:"username"`
	Gender   string      `json:"gender"`
	Security SecurityDTO `json:"security"`
}

type UpdateRequestDTO struct {
	Id       string      `json:"id"`
	Email    string      `json:"email"`
	Role     string      `json:"role"`
	Username string      `json:"username"`
	Gender   string      `json:"gender"`
	Security SecurityDTO `json:"security"`
}

type DeleteRequestDTO struct {
	Id string `json:"id"`
}
