package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from GetAllUsers!")
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from GetUserByID!")
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from CreateUser!")
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from DeleteUser!")
}
