package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	tenPercentCallback         = "menuCallback"
	twentyFivePercentCallback  = "callbackPlay"
	fiftyPercentCallback       = "callbackBalance"
	seventyFivePercentCallback = "callbackRefill"
	maxPercentCallback         = "callbackWithdraw"
)

func generateBalanceButtons(balance int) []tgbotapi.InlineKeyboardButton {
	var output []tgbotapi.InlineKeyboardButton

	if balance == 0 {
		return output
	}

	if balance <= 10 {
		output = append(output, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", balance), maxPercentCallback))
		return output
	}

	output = append(output, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", int(float64(balance)*0.1)), tenPercentCallback))
	output = append(output, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", int(float64(balance)*0.25)), twentyFivePercentCallback))
	output = append(output, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", int(float64(balance)*0.5)), fiftyPercentCallback))
	output = append(output, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", int(float64(balance)*0.75)), seventyFivePercentCallback))
	output = append(output, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", balance), maxPercentCallback))

	return output
}
