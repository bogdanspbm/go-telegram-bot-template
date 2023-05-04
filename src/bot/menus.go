package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-telegram-bot-template/src/utils"
)

func (bot *bot) getStartMenu(msgID int, chatID int64) tgbotapi.EditMessageTextConfig {
	btnPlay := tgbotapi.NewInlineKeyboardButtonData("Играть", playCallback)
	btnBalance := tgbotapi.NewInlineKeyboardButtonData("Баланс", balanceCallback)
	btnRefill := tgbotapi.NewInlineKeyboardButtonData("Пополнить", refillCallback)
	btnWithdraw := tgbotapi.NewInlineKeyboardButtonData("Снять", withdrawCallback)

	rowTop := tgbotapi.NewInlineKeyboardRow(btnPlay, btnBalance)
	rowBot := tgbotapi.NewInlineKeyboardRow(btnRefill, btnWithdraw)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rowTop, rowBot)

	message := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, "Выберите действие: ", keyboard)
	return message
}

func (bot *bot) getBalanceMenu(msgID int, chatID int64) tgbotapi.EditMessageTextConfig {
	btnBack := tgbotapi.NewInlineKeyboardButtonData("Назад", menuCallback)
	btnRefresh := tgbotapi.NewInlineKeyboardButtonData("Обновить", balanceCallback)

	row := tgbotapi.NewInlineKeyboardRow(btnBack, btnRefresh)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	body := fmt.Sprintf("Ваш баланс: %v р.", bot.storage.getBalance(chatID))
	message := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, body, keyboard)
	return message
}

func (bot *bot) getWithdrawMenu(msgID int, chatID int64) tgbotapi.EditMessageTextConfig {
	btnBack := tgbotapi.NewInlineKeyboardButtonData("Назад", menuCallback)

	buttons := generateBalanceButtons(bot.storage.getBalance(chatID))
	buttons = append(buttons, btnBack)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(utils.ConvertButtonsToRows(2, buttons...)...)

	body := fmt.Sprintf("Ваш баланс: %v р. \nКакую сумму вы желаете вывести?", bot.storage.getBalance(chatID))

	if len(buttons) == 0 {
		body = fmt.Sprintf("Ваш баланс: 0 р. \nПополните счет, чтобы вывести.", bot.storage.getBalance(chatID))
	}

	if len(buttons) == 1 {
		body = fmt.Sprintf("Ваш баланс: %v р. \nПодтвердите вывод всей суммы.", bot.storage.getBalance(chatID))
	}

	message := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, body, keyboard)
	return message
}
