package datatransfer

type AdminRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
