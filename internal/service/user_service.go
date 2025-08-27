// Package service
package service

import (
	"context"
	"fmt"
	"os"
	"time"
	"user-management-service/internal/entity"
	"user-management-service/internal/repository"

	//"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

type UserService struct {
	userRepo repository.UserRepository
	rdb      *redis.Client
}
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
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
	return user, nil
}

func (u *UserService) CreateUser(user *entity.User) (*entity.User, error) {
	createdUser, err := u.userRepo.CreateUser(user)
	if err != nil {
		logger.Error().Err(err).Msgf("Error creating user")
		return nil, err
	}
	return createdUser, nil
}

func (u *UserService) Login(ctx context.Context, email, password string) (token string, err error) {
	user, err := u.userRepo.GetUserByEmailAndPassword(email, password)
	// validate the user credentials against a database
	if err != nil {
		return "", err
	}

	// after validation, generate JWT token
	claims := &JwtCustomClaims{
		Name:  user.Username,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := tkn.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	// store jwt in redis with the user email as the key
	err = u.rdb.Set(ctx, email, t, time.Hour*24).Err()
	if err != nil {
		return "", err
	}

	// user , err := u.userRepo.GetUserByEmailAndPassword(email, password)
	// if err != nil {
	// 	logger.Error().Err(err).Msgf("Error loggin in user")
	// 	return nil, err
	// }
	return t, nil

}

func (u *UserService) ValidateToken(ctx context.Context, email string) (string, error) {
	// retrieve the jwtt token from redis
	token, err := u.rdb.Get(ctx, email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("session not found")
		}
		return "", err
	}
	return token, nil
}
