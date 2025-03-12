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

type UserByIDRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type UserByIDResponse struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserByEmailRequest struct {
	Email string `json:"email" param:"email" validate:"required"`
}

type UserByEmailResponse struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	ID       string `json:"id" param:"id" validate:"required"`
	Email    string `json:"email" param:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserUpdateResponse struct {
	Email    *string `json:"email,omitempty"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
}
