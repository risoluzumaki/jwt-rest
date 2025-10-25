package auth

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUserDTO struct {
	Token string `json:"token"`
}
