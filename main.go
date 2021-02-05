package main

import (
	"os"

	"gobot/chatbot"
)

func main() {
	port := os.Getenv("PORT")
	bot := chatbot.New(port)
	bot.Start()
}
