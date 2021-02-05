package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	listeningPort := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(fmt.Sprint(":%s", listeningPort)))
}
