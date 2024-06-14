package utils

import (
	"github.com/mymmrac/telego"
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
