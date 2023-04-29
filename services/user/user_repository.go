package user

import "golang-web-testing/config"

type UserRepositoryInterface interface {
	Create(UserRequest) (User, error)
	FindByEmail(string) (User, bool)
}

type UserRepository struct{}

func (repository UserRepository) Create(u UserRequest) (User, error) {
	user := User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repository UserRepository) FindByEmail(email string) (user User, found bool) {
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, false
	}

	return user, true
}
