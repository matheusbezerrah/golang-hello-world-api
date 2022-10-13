package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", hello)

	e.Logger.Fatal(e.Start(":1532"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hellow World")
}
