package chatbot

import (
	"fmt"
	"gobot/models"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const defaultPort = "8080"

// Chatbot struct
type Chatbot struct {
	Port string
}

// New a Chatbot
func New(port string) *Chatbot {
	if port == "" {
		port = defaultPort
	}
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

	b.showBanner()
	e.HideBanner = true

	e.Logger.Info("Migrating database if needs ...")
	b.migrateDB()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", b.Port)))
}

func (b *Chatbot) migrateDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(pq.Open(databaseURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Phonebook{})
}

func (b *Chatbot) showBanner() {
	color.Println(color.Blue(`
 ██████╗  ██████╗ ██████╗  ██████╗ ████████╗
██╔════╝ ██╔═══██╗██╔══██╗██╔═══██╗╚══██╔══╝
██║  ███╗██║   ██║██████╔╝██║   ██║   ██║   
██║   ██║██║   ██║██╔══██╗██║   ██║   ██║   
╚██████╔╝╚██████╔╝██████╔╝╚██████╔╝   ██║   
 ╚═════╝  ╚═════╝ ╚═════╝  ╚═════╝    ╚═╝   `))
}
