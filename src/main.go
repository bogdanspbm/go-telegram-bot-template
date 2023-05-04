package main

import bot "go-telegram-bot-template/src/bot"

func main() {
	tgBot, err := bot.New("TOKEN")

	if err != nil {
		panic("can't create bot")
	}

	tgBot.Start(true)
}
