package goclicker

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// StartServer is
func StartServer() {
	db := GetDatabase()
	db.AutoMigrate(&Link{}, &AccessLog{})
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.GET("/api/links", GetLinksHandler, CheckAuth)
	e.GET("/api/links/:key", GetLinkHandler, CheckAuth)
	e.POST("/api/links", SaveLinkHandler, CheckAuth)
	e.PATCH("/api/links/:id", UpdateLinkHandler, CheckAuth)
	e.GET("/:key", LinkHandler)
	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
