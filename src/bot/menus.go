package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	playCallback     = "callbackPlay"
	balanceCallback  = "callbackBalance"
	refillCallback   = "callbackRefill"
	withdrawCallback = "callbackWithdraw"
)

func getStartMenu(chatID int64) tgbotapi.MessageConfig {
	btnPlay := tgbotapi.NewInlineKeyboardButtonData("Играть", playCallback)
	btnBalance := tgbotapi.NewInlineKeyboardButtonData("Баланс", balanceCallback)
	btnRefill := tgbotapi.NewInlineKeyboardButtonData("Пополнить", refillCallback)
	btnWithdraw := tgbotapi.NewInlineKeyboardButtonData("Снять", withdrawCallback)

	rowTop := tgbotapi.NewInlineKeyboardRow(btnPlay, btnBalance)
	rowBot := tgbotapi.NewInlineKeyboardRow(btnRefill, btnWithdraw)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rowTop, rowBot)

	message := tgbotapi.NewMessage(chatID, "Выберите действие:")
	message.ReplyMarkup = keyboard
	return message
}
