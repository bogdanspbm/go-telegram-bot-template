package bot

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type bot struct {
	api     *telegram.BotAPI
	storage *storage
}

func New(token string) (*bot, error) {
	api, err := telegram.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	tgBot := &bot{api: api, storage: createStorage()}
	return tgBot, nil
}

func (bot *bot) Start(debug bool) {
	bot.api.Debug = debug

	log.Printf("Authorized on account %s", bot.api.Self.UserName)

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates := bot.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go bot.handleMessage(*update.Message)
		}
		if update.CallbackQuery != nil {
			go bot.handleCallback(*update.CallbackQuery)
		}
	}
}

func (bot *bot) handleMessage(message telegram.Message) error {

	if message.IsCommand() {
		return bot.handleCommand(message)
	}

	msg := telegram.NewMessage(message.Chat.ID, "Слушаю:")
	msg.ReplyToMessageID = message.MessageID
	bot.api.Send(msg)
	return nil
}
