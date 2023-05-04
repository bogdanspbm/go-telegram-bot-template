package bot

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	tenPercentCallback         = "tenPercentCallback"
	twentyFivePercentCallback  = "twentyFivePercentCallback"
	fiftyPercentCallback       = "fiftyPercentCallback"
	seventyFivePercentCallback = "seventyFivePercentCallback"
	maxPercentCallback         = "maxPercentCallback"
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

func (bot *bot) handleWithdrawCallback(query tgbotapi.CallbackQuery) error {
	balance := bot.storage.getBalance(query.Message.Chat.ID)

	switch query.Data {
	case tenPercentCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, -int(float64(balance)*0.1))
			break
		}
	case twentyFivePercentCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, -int(float64(balance)*0.25))
			break
		}
	case fiftyPercentCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, -int(float64(balance)*0.5))
			break
		}
	case seventyFivePercentCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, -int(float64(balance)*0.75))
			break
		}
	case maxPercentCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, -balance)
			break
		}
	default:
		return errors.New("isn't withdraw callback")
	}

	bot.api.Send(bot.getWithdrawMenu(query.Message.MessageID, query.Message.Chat.ID))
	return nil
}
