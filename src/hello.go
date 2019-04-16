package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/env", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"SECRET_VALUE": os.Getenv("SECRET_VALUE"),
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
