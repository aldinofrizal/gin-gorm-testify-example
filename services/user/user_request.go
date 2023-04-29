package user

type UserRequest struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
