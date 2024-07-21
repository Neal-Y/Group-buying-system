package forgot_pwd

type GetEmailByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}
