package main

import (
	"time"
	"user-management-service/internal/api"
	"user-management-service/internal/repository"
	"user-management-service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)
func main(){
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(*userRepo)
	userHandler := api.NewUserHandler(*userService)

	e := echo.New()

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate: rate.Limit(1),
				Burst: 3,
				ExpiresIn: 3 * time.Minute,
			}),
		IdentifierExtractor: func(context echo.Context) (string, error) {
			// for local
			return context.Request().RemoteAddr, nil
			// for production
			// return context.Request().Header.Get(echo.HeaderXRealIP), nil
			//return context.RealIP(), nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(429, map[string]string{"error":"rate limit exceed"})
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(429, map[string]string{"error":"rate limit exceed"})
		},
	}
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiterWithConfig(config))
	//e.Use(echojwt.JWT([]byte("secret")))

	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.CreateUser)
	e.POST("/login", userHandler.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
