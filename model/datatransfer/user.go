package datatransfer

type LineProfileResponse struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type LineTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRequest struct {
	LineID      string `json:"line_id"`
	DisplayName string `json:"display_name" binding:"required"`
	Email       string `json:"email"`
	LineToken   string `json:"line_token"`
	Phone       string `json:"phone"`
	IsMember    bool   `json:"is_member"`
}

type UserResponse struct {
	ID          int    `json:"id"`
	LineID      string `json:"line_id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	LineToken   string `json:"line_token"`
	Phone       string `json:"phone"`
	IsMember    bool   `json:"is_member"`
}
