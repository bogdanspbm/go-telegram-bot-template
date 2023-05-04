package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-telegram-bot-template/src/utils"
)

const (
	menuCallback = "callbackMenu"

	betBlackCallback = "callbackBetBlack"
	betRedCallback   = "callbackBetRed"

	changeOneCallback     = "callbackChangeOne"
	changeTenCallback     = "callbackChangeTen"
	changeHundredCallback = "callbackChangeHundred"
	playCallback          = "callbackPlay"
	changeBetCallback     = "callbackChangeBet"
	balanceCallback       = "callbackBalance"
	refillCallback        = "callbackRefill"
	withdrawCallback      = "callbackWithdraw"
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
	case playCallback:
		{
			bot.api.Send(bot.getPlayMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case betRedCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, utils.RandSign()*bot.storage.getBet(query.Message.Chat.ID))
			bot.api.Send(bot.getPlayMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case betBlackCallback:
		{
			bot.storage.addBalance(query.Message.Chat.ID, utils.RandSign()*bot.storage.getBet(query.Message.Chat.ID))
			bot.api.Send(bot.getPlayMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}

	case changeOneCallback:
		{
			bot.storage.setBet(query.Message.Chat.ID, 1)
			bot.api.Send(bot.getPlayMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case changeTenCallback:
		{
			bot.storage.setBet(query.Message.Chat.ID, 10)
			bot.api.Send(bot.getPlayMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case changeHundredCallback:
		{
			bot.storage.setBet(query.Message.Chat.ID, 100)
			bot.api.Send(bot.getPlayMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}
	case changeBetCallback:
		{
			bot.api.Send(bot.getChangeBetMenu(query.Message.MessageID, query.Message.Chat.ID))
			return
		}

	default:
		bot.handleWithdrawCallback(query)
	}

}
