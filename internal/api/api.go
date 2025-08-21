package api

import (
	"strconv"
	"time"
	"user-management-service/internal/entity"
	"user-management-service/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService)*UserHandler{
	return &UserHandler{
		userService: userService,
	}
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (u *UserHandler)GetUserByID(c echo.Context)error{
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]string{"error":"Invalid ID"})
	}
	user, err := u.userService.GetUserByID(idInt)
	if err != nil {
		return c.JSON(500, map[string]string{"error":err.Error()})
	}
	return c.JSON(200, user)
}

func (u *UserHandler) CreateUser(c echo.Context)error {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]string{"error":"Invalid request payload"})
	}

	createdUser, err := u.userService.CreateUser(&user)
	if err != nil {
		return c.JSON(500, map[string]string{"error":err.Error()})
	}
	return c.JSON(200, createdUser)
}

// login
func (u *UserHandler)Login(c echo.Context)error{
	login := struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&login);err != nil {
		return c.JSON(400, map[string]string{"error":"Invalid request payload"})
	}
	user, err := u.userService.Login(login.Email,login.Password)
	if err != nil {
		return c.JSON(401, map[string]string{"error":err.Error()})
	}

	if user == nil {
		return c.JSON(401, map[string]string{"error":"invalid email or password"})
	}

	claims := &JwtCustomClaims{
		Name : user.Username,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.JSON(401, map[string]string{"error":err.Error()})
	}
	return c.JSON(200, map[string]string{"token":t})	
}