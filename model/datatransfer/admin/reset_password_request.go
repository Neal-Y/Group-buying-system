package admin

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
