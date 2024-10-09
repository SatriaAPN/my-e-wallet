package core

type RestCreateUserRequest struct {
	Name     string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RestCreateUserResponse struct {
	Name  string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type RestLoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RestLoginUserResponse struct {
	Token string `json:"token"`
}
