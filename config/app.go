package config

import (
	"database/sql"
	db "github.com/LuccChagas/clean-scaffold/db/sqlc"
	"github.com/LuccChagas/clean-scaffold/internal/app/handler"
	"github.com/LuccChagas/clean-scaffold/internal/app/repository"
	"github.com/LuccChagas/clean-scaffold/internal/app/router"
	"github.com/LuccChagas/clean-scaffold/internal/app/service"
)

type RepositoryInstance struct {
	Repository *repository.Repository
}

type ServiceInstance struct {
	UserService *service.UserService
}

type HandlerInstance struct {
	UserHandler *handler.UserHandler
}

func newRepositoryInstance(sqlDB *sql.DB) *RepositoryInstance {
	repoInstance := &RepositoryInstance{
		Repository: repository.NewRepository(sqlDB, db.New(sqlDB)),
	}

	return repoInstance
}

func newServiceInstance(repoInstance *RepositoryInstance) *ServiceInstance {
	serviceInstance := &ServiceInstance{
		UserService: service.NewUserService(repoInstance.Repository),
	}
	return serviceInstance
}

func newHandlerInstance(serviceInstance *ServiceInstance) *HandlerInstance {
	handlerInstance := &HandlerInstance{
		UserHandler: handler.NewUserHandler(serviceInstance.UserService),
	}
	return handlerInstance
}

type App struct {
	Server *router.Router
}

func NewApp(db *sql.DB) *App {

	repoInstance := newRepositoryInstance(db)
	serviceInstance := newServiceInstance(repoInstance)
	handlerInstance := newHandlerInstance(serviceInstance)

	server := router.NewRouter(
		handlerInstance.UserHandler)

	return &App{
		Server: server,
	}
}
