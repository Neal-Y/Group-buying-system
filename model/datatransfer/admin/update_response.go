package admin

type UpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	LineID   string `json:"line_id"`
}
