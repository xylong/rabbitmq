package Models

type UserModel struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name" binding:"required"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}
