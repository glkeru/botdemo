package commands

import (
	"botmodul/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandRouter struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(tgbot *tgbotapi.BotAPI, productService *product.Service) *CommandRouter {
	return &CommandRouter{
		bot:            tgbot,
		productService: productService,
	}
}

func (c *CommandRouter) RunCommand(inputmsg *tgbotapi.Message) {

	switch inputmsg.Command() {
	case "help":
		c.helpCommand(inputmsg)
	case "list":
		c.listCommand(inputmsg)
	default:
		c.defaultBehavior(inputmsg)
	}

}
