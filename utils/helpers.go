package utils

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
)

// MustOpen - Возвращает ссылку на открытый / созданный файл
func MustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

// DelMessage - Удаляем сообщение, которое расположено перед вызовом этой функции
func DelMessage(bot *telego.Bot, update telego.Update) {
	mess := update.Message
	var DelMessParams telego.DeleteMessageParams
	DelMessParams.ChatID = mess.Chat.ChatID()
	DelMessParams.MessageID = mess.MessageID
	_ = bot.DeleteMessage(&DelMessParams)
}

func InlineKeyboardButtonSend(nameButton string, nameCallback string) *telego.InlineKeyboardMarkup {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(nameButton).
				WithCallbackData(nameCallback),
		),
	)
	return inlineKeyboard
}

func ConfirmSend(nameCallback string) *telego.InlineKeyboardMarkup {
	return InlineKeyboardButtonSend("✔️Подтвердить отправку", nameCallback)
}
