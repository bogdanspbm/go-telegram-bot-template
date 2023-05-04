package main

import bot "go-telegram-bot-template/src/bot"

func main() {
	tgBot, err := bot.New("6142838792:AAHCYvB7JNLeTJXwEecNheAsgZCUu6Df_Ys")

	if err != nil {
		panic("can't create bot")
	}

	tgBot.Start(true)
}
