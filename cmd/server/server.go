package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Cogitator activated.")
	})

	e.Logger.Fatal(e.Start(":8080"))

}
