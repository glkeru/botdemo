package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *CommandRouter) helpCommand(inputmsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputmsg.Chat.ID, "HELP!!")
	msg.ReplyToMessageID = inputmsg.MessageID
	c.bot.Send(msg)
}
