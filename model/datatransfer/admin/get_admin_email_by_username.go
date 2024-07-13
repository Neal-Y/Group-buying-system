package admin

type GetAdminEmailByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}
