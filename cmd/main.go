package main

import (
	"user-management-service/internal/api"
	"user-management-service/internal/repository"
	"user-management-service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func main(){
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(*userRepo)
	userHandler := api.NewUserHandler(*userService)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Use(echojwt.JWT([]byte("secret")))

	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
