package main

import (
	"user-management-service/internal/api"
	"user-management-service/internal/service"

	"github.com/labstack/echo/v4"
)
func main(){
	userService := service.NewUserService()
	userHandler := api.NewUserHandler(*userService)

	e := echo.New()

	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)

	e.Logger.Fatal(e.Start(":8088"))
}
