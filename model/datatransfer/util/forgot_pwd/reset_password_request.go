package forgot_pwd

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
