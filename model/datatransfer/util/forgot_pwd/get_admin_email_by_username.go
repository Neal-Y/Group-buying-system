package forgot_pwd

type GetAdminEmailByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}
