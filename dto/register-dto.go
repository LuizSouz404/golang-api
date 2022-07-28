package dto

// RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name     string `json:"name" from:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password" from:"password" validate:"min:6" binding:"required"`
}
