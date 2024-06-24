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

var MessIdForDel int = 0
var MessFlagForDel bool = false
var CurUpdate telego.Update // Задействуется для вызова функции из роутера сайта

// DelMessage - Удаляем сообщение, которое расположено перед вызовом этой функции
// Если MessIdForDel != 0 и MessFlagForDel == true, удаляется сообщение с id в переменной MessIdForDel
func DelMessage(bot *telego.Bot, update telego.Update) {
	mess := update.Message
	var DelMessParams telego.DeleteMessageParams
	DelMessParams.ChatID = mess.Chat.ChatID()
	DelMessParams.MessageID = mess.MessageID
	if MessIdForDel != 0 && MessFlagForDel == true {
		DelMessParams.MessageID = MessIdForDel
		MessIdForDel = 0
		MessFlagForDel = false
	}
	_ = bot.DeleteMessage(&DelMessParams)
}

func DelMessagesByName(mess string, bot *telego.Bot, update ...telego.Update) string {
	//DelMessage()

	return "777777777777777"
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
