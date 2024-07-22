package forgot_pwd

type NewPasswordRequest struct {
	Email       string `json:"email" binding:"required"`
	Code        string `json:"code" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	Username    string `json:"username" binding:"required"`
}
