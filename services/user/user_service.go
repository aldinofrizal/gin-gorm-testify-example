package user

import "errors"

type UserService struct {
	Repository UserRepositoryInterface
}

func (service UserService) Register(userReq UserRequest) (User, error) {
	var user User

	_, found := service.Repository.FindByEmail(userReq.Email)
	if found {
		return User{}, errors.New("email already in use")
	}

	user, err := service.Repository.Create(userReq)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (service UserService) Login(email, password string) (User, error) {
	var user User
	user, found := service.Repository.FindByEmail(email)
	if !found {
		return user, errors.New("email/password invalid")
	}

	if user.Password != password {
		return user, errors.New("email/password invalid")
	}

	return user, nil
}
