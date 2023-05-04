package utils

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
)

func ConvertButtonsToRows(columns int, buttons ...telegram.InlineKeyboardButton) (rows [][]telegram.InlineKeyboardButton) {

	row := make([]telegram.InlineKeyboardButton, 0)

	for i := 0; i < len(buttons); i++ {
		var button = buttons[i]
		row = append(row, button)

		if len(row) == columns {
			rows = append(rows, row)
			row = make([]telegram.InlineKeyboardButton, 0)
		}
	}

	if len(rows) == 0 {
		rows = append(rows, row)
	}

	return rows
}

func RandSign() int {
	return rand.Intn(2)*2 - 1
}
