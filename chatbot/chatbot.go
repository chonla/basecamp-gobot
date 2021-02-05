package chatbot

import (
	"fmt"
	"gobot/models"
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
	e.POST("/", func(c echo.Context) error {
		m := new(models.MessageRequest)
		if err := c.Bind(m); err != nil {
			return err
		}

		response := fmt.Sprintf("พูดอะไรนะ <b>%s</b>", m.Command)
		return c.String(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", b.Port)))
}
