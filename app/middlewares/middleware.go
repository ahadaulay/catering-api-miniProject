package middlewares

import (
	middlewares "catering-api/app/middlewares/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigLogger struct {
	Format string
}

func (c *ConfigLogger) Init() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: c.Format,
	})
}

func CheckTokenMiddlewareUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.GetUserCustomer(c)

		if userID == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"messege": "invalid create token",
			})
		}
		return next(c)
	}
}
