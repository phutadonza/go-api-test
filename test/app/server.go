package main

import (
	"test/app/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userHandler := handlers.NewUserHandler()

	e.GET("/users", userHandler.GetAllUsers)

	e.GET("/users/:id", userHandler.GetUserByID)

	e.POST("/users", userHandler.CreateUser)

	e.DELETE("/users/:id", userHandler.DeleteUser)

	e.Start(":8000")
}
