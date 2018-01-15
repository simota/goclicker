package goclicker

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// CheckAuth is API authentication middleware
func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-ApiKey")
		if os.Getenv("API_KEY") != apiKey {
			return c.String(http.StatusUnauthorized, "required api key")
		}
		return next(c)
	}
}
