package auth

type LoginRequest struct {
	Email    string `json:"email" form:"email" xml:"email" binding:"required,email"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}

type SignupRequest struct {
	Email    string `json:"email" form:"email" xml:"email" binding:"required,email"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
	Name     string `json:"name" form:"name" xml:"name" binding:"required"`
}
