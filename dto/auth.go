package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID          int    `json:"id"`
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,password"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Role            string `json:"role" binding:"required,oneof='admin' 'staff' 'guest'"`
}
