package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (router *Router) endpoints(e *echo.Echo) {

	user := e.Group("/user")
	user.POST("/create", router.User.CreateUserHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
