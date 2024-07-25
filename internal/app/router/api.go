package router

import (
	"github.com/LuccChagas/clean-scaffold/internal/app/handler"
	"github.com/labstack/echo/v4"
)

// Router - Add here all handler interfaces
type Router struct {
	User handler.UserHandlerInterface
}

func NewRouter(user handler.UserHandlerInterface) *Router {
	return &Router{
		User: user,
	}
}

// Serve
// @title Swagger project API
// @version 1.0
// @description Document API
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host www
// @BasePath /
// @schemes https
func (router *Router) Serve() {
	e := echo.New()
	loadMiddlewares(e)

	router.endpoints(e)

	e.Logger.Fatal(e.Start(":8080"))

}
