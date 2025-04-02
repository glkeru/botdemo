package main

import (
	"log"
	"os"

	"botmodul/internal/app/commands"
	"botmodul/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOTTOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		commander.RunCommand(update.Message)

	}
}
