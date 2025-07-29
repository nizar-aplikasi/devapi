package dto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	NoTelp   string `json:"notelp" binding:"required"`
	OrgName  string `json:"orgname"`
	Role     string `json:"role"`
}
