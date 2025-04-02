package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *CommandRouter) listCommand(inputmsg *tgbotapi.Message) {
	listMsg := "List: \n"
	products := c.productService.List()
	for _, p := range products {
		listMsg += "\n" + p.Title
	}

	msg := tgbotapi.NewMessage(inputmsg.Chat.ID, listMsg)
	msg.ReplyToMessageID = inputmsg.MessageID
	c.bot.Send(msg)
}
