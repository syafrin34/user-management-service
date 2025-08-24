package repository

import (
	"user-management-service/internal/entity"
)

type UserRepository struct {

}

func NewUserRepository()*UserRepository{
	return  &UserRepository{}
}

// mock user db

var users = map[int]*entity.User{
	1:{ID: 1, Username: "user1", Email: "user1@gmail.com",Password: "12345"},
	2:{ID: 2, Username: "user2", Email: "user2@gmail.com",Password: "12345"},
}
func (u *UserRepository)GetUserByID(id int)(*entity.User, error){
	user, ok := users[id]
	if !ok {
		return  nil, nil
	}
	return user, nil
}
func (u *UserRepository)GetUserByEmail(email string)(*entity.User, error){
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, nil
}
func (u *UserRepository)CreateUser(user *entity.User)(*entity.User, error){
	user.ID = 3
	users[user.ID] = user
	return  user, nil

}

func (u *UserRepository)GetUserByEmailAndPassword(email, password string)(*entity.User, error){
	for _, user := range users{
		if user.Email == email && user.Password == password {
			return  user, nil
		}
	}

	return nil, nil
}
