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
	msg := telegram.NewMessage(message.Chat.ID, fmt.Sprintf("Привет, %v!", message.From.FirstName))
	bot.api.Send(msg)
	bot.api.Send(getStartMenu(message.Chat.ID))
	return nil
}
