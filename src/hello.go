package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
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
	e.GET("/health_db", dbHandler)
	e.GET("/health_redis", redisHandler)
	e.GET("/env", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"SECRET_VALUE": os.Getenv("SECRET_VALUE"),
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func dbHandler(c echo.Context) error {
	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, version)
}

func redisHandler(c echo.Context) error {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	result, err := client.Info("server").Result()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, result)
}
