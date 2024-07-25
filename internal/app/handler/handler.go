package handler

import "github.com/labstack/echo/v4"

type UserHandlerInterface interface {
	CreateUserHandler(c echo.Context) error
}
