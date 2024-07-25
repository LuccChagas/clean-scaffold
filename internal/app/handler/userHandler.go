package handler

import (
	"github.com/LuccChagas/clean-scaffold/internal/app/model"
	"github.com/LuccChagas/clean-scaffold/internal/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(u service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: u,
	}
}

func (h *UserHandler) CreateUserHandler(c echo.Context) error {
	var request model.User
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.service.CreateUserService(c.Request().Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response)
}
