package service

import (
	"os"
	"user-management-service/internal/entity"
	"user-management-service/internal/repository"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetUserByID(idUser int) (*entity.User, error) {
	user, err := u.userRepo.GetUserByID(idUser)
	if err != nil {
		logger.Error().Err(err).Msgf("Error getting userby ID %d", idUser)
		return nil, err
	}
	return  user, nil
}

func (u *UserService) CreateUser(user *entity.User) (*entity.User, error) {
	createdUser , err := u.userRepo.CreateUser(user)
	if err != nil {
		logger.Error().Err(err).Msgf("Error creating user")
		return  nil , err
	}
	return createdUser, nil
}

func (u *UserService)Login(email, password string)(*entity.User, error){
	user , err := u.userRepo.GetUserByEmailAndPassword(email, password)
	if err != nil {
		logger.Error().Err(err).Msgf("Error loggin in user")
		return nil, err	
	}
	return user, nil

}


