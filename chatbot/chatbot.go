package chatbot

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Chatbot struct
type Chatbot struct {
	Port string
}

// New a Chatbot
func New(port string) *Chatbot {
	return &Chatbot{
		Port: port,
	}
}

// Start the server
func (b *Chatbot) Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", b.Port)))
}
