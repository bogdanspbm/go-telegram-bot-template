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

func (bot *bot) getPlayMenu(msgID int, chatID int64) tgbotapi.EditMessageTextConfig {
	btnBack := tgbotapi.NewInlineKeyboardButtonData("Назад", menuCallback)
	btnChange := tgbotapi.NewInlineKeyboardButtonData("Поменять ставку", changeBetCallback)
	btnBlack := tgbotapi.NewInlineKeyboardButtonData("Черное", betBlackCallback)
	btnRed := tgbotapi.NewInlineKeyboardButtonData("Красное", betRedCallback)

	rowTop := tgbotapi.NewInlineKeyboardRow(btnBlack, btnRed)
	rowDown := tgbotapi.NewInlineKeyboardRow(btnBack, btnChange)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rowTop, rowDown)

	body := fmt.Sprintf("Ваш баланс: %v р. \nТекущая ставка: %v р.\n Выберите на что хотите поставить:", bot.storage.getBalance(chatID), bot.storage.getBet(chatID))
	message := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, body, keyboard)
	return message
}

func (bot *bot) getChangeBetMenu(msgID int, chatID int64) tgbotapi.EditMessageTextConfig {
	btnBack := tgbotapi.NewInlineKeyboardButtonData("Назад", playCallback)
	btnOne := tgbotapi.NewInlineKeyboardButtonData("1", changeOneCallback)
	btnTen := tgbotapi.NewInlineKeyboardButtonData("10", changeTenCallback)
	btnHundred := tgbotapi.NewInlineKeyboardButtonData("100", changeHundredCallback)

	rowTop := tgbotapi.NewInlineKeyboardRow(btnOne, btnTen)
	rowDown := tgbotapi.NewInlineKeyboardRow(btnBack, btnHundred)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rowTop, rowDown)

	body := fmt.Sprintf("Выберите сумму ставки:")
	message := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, body, keyboard)
	return message
}

func (bot *bot) getWithdrawMenu(msgID int, chatID int64) tgbotapi.EditMessageTextConfig {
	btnBack := tgbotapi.NewInlineKeyboardButtonData("Назад", menuCallback)

	buttons := generateBalanceButtons(bot.storage.getBalance(chatID))
	buttons = append(buttons, btnBack)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(utils.ConvertButtonsToRows(2, buttons...)...)

	body := fmt.Sprintf("Ваш баланс: %v р. \nКакую сумму вы желаете вывести?", bot.storage.getBalance(chatID))

	if len(buttons) == 1 {
		body = fmt.Sprintf("Ваш баланс: 0 р. \nПополните счет, чтобы вывести.")
	}

	if len(buttons) == 2 {
		body = fmt.Sprintf("Ваш баланс: %v р. \nПодтвердите вывод всей суммы.", bot.storage.getBalance(chatID))
	}

	message := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, body, keyboard)
	return message
}
