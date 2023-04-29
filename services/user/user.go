package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Name     string
	Password string
}

func (u *User) GetBasicResponse() UserResponse {
	return UserResponse{
		Id:    int(u.ID),
		Email: u.Email,
		Name:  u.Name,
	}
}
