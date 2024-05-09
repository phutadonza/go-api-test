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
	return c.HTML(http.StatusOK, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>GetAllUsers</title>
            <style>
                body {
                    background-color: #f44336; /* สีแดง */
                    color: white;
                }
            </style>
        </head>
        <body>
            <h1>Hello from GetAllUsers</h1>
        </body>
        </html>
    `)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	return c.HTML(http.StatusOK, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>GetUserByID</title>
            <style>
                body {
                    background-color: #2196F3; /* สีน้ำเงิน */
                    color: white;
                }
            </style>
        </head>
        <body>
            <h1>Hello from GetUserByID!</h1>
        </body>
        </html>
    `)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	return c.HTML(http.StatusOK, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>CreateUser</title>
            <style>
                body {
                    background-color: #4CAF50; /* สีเขียว */
                    color: white;
                }
            </style>
        </head>
        <body>
            <h1>Hello from CreateUser!</h1>
        </body>
        </html>
    `)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	return c.HTML(http.StatusOK, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>DeleteUser</title>
            <style>
                body {
                    background-color: #FFC107; /* สีส้ม */
                    color: white;
                }
            </style>
        </head>
        <body>
            <h1>Hello from DeleteUser!</h1>
        </body>
        </html>
    `)
}
