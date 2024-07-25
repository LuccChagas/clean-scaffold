package router

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/LuccChagas/clean-scaffold/internal/pkg/token"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func loadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, //temp
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: middleware.DefaultCORSConfig.AllowMethods,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			e.Logger.Info("request",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.Int("status", c.Response().Status),
				zap.Duration("duration", stop.Sub(start)),
			)

			return err
		}
	})
}

func checkAuthorization(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		bearerToken := c.Request().Header.Get("Authorization")
		tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

		maker, err := token.NewPasetoMaker(os.Getenv("TOKEN_SIGNATURE"))
		if err != nil {
			return c.JSON(http.StatusBadGateway, err.Error())
		}

		tokenPayload, err := maker.VerifyToken(tokenStr)
		if err != nil {
			return c.JSON(http.StatusBadGateway, err.Error())
		}
		c.Set("token_id", tokenPayload.ID)
		c.Set("token_user_id", tokenPayload.UserID)
		c.Set("token_user_name", tokenPayload.Username)
		c.Set("token_access_key", tokenPayload.AccessKey)
		c.Set("token_tenant_id", tokenPayload.TenantID)
		c.Set("token_expiry_at", tokenPayload.ExpiredAt)

		return handlerFunc(c)
	}
}
