package service

import "user-management-service/internal/entity"

type UserService struct {
	// reposiroty
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserByID(idUser int) (*entity.User, error) {
	user := &entity.User{
		ID:       idUser,
		Username: "Test User",
		Email:    "user@gmail.com",
		Password: "1234",
	}
	return user, nil
}

func (u *UserService) CreateUser(user *entity.User) (*entity.User, error) {
	user.ID = 1
	return user, nil
}

func (u *UserService)Login(email, password string)(*entity.User, error){
	user := &entity.User{
		ID: 1,
		Username: "user",
		Email: email,
	}
	return  user, nil
}


