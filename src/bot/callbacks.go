package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	menuCallback     = "menuCallback"
	playCallback     = "callbackPlay"
	balanceCallback  = "callbackBalance"
	refillCallback   = "callbackRefill"
	withdrawCallback = "callbackWithdraw"
)

func (bot *bot) handleCallback(query tgbotapi.CallbackQuery) {
	switch query.Data {
	case menuCallback:
		{
			bot.api.Send(bot.getStartMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case balanceCallback:
		{
			bot.api.Send(bot.getBalanceMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case withdrawCallback:
		{
			bot.api.Send(bot.getWithdrawMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	}

}
