package user

type Register struct {
	DisplayName string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone"`
}
