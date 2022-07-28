package dto

// LoginDTO is used when client post from /login url
type LoginDTO struct {
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password" from:"password" validate:"min:6" binding:"required"`
}
