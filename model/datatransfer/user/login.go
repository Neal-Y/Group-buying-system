package user

type Login struct {
	DisplayName string `json:"display_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
