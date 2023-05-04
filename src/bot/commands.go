package bot

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *bot) handleCommand(message telegram.Message) error {

	switch message.Command() {
	case "start":
		{
			return bot.handleStartCommand(message)
		}
	}

	return nil
}

func (bot *bot) handleStartCommand(message telegram.Message) error {
	msgConf := telegram.NewMessage(message.Chat.ID, fmt.Sprintf("Привет, %v!", message.From.FirstName))
	bot.api.Send(msgConf)

	menuMsgConf := telegram.NewMessage(message.Chat.ID, "Выберите действие: ")
	msg, err := bot.api.Send(menuMsgConf)

	if err != nil {
		return err
	}

	bot.api.Send(bot.getStartMenu(msg.MessageID, msg.Chat.ID))
	return nil
}
