package dto

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserByEmailRequest struct {
	Email string `json:"email" param:"email" validate:"required"`
}

type UserByEmailResponse struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
