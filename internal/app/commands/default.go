package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *CommandRouter) defaultBehavior(inputmsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputmsg.Chat.ID, "You wrote: "+inputmsg.Text)
	msg.ReplyToMessageID = inputmsg.MessageID
	c.bot.Send(msg)
}
