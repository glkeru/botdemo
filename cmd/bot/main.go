package main

import (
	"log"
	"os"

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

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		default:
			defaultBehavior(bot, update.Message)
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputmsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputmsg.Chat.ID, "HELP!!")
	msg.ReplyToMessageID = inputmsg.MessageID
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputmsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputmsg.Chat.ID, "You wrote: "+inputmsg.Text)
	msg.ReplyToMessageID = inputmsg.MessageID
	bot.Send(msg)
}
