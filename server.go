package main

import (
	"net/http"
	 db "github.com/KepedulianMasyarakat/go-kemas/config"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}