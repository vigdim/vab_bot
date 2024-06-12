package utils

import (
	"github.com/mymmrac/telego"
	"os"
)

// Helper function to open file or panic
func MustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func DelMessage(bot *telego.Bot, update telego.Update) {
	mess := update.Message
	var DelMessParams telego.DeleteMessageParams
	DelMessParams.ChatID = mess.Chat.ChatID()
	DelMessParams.MessageID = mess.MessageID
	_ = bot.DeleteMessage(&DelMessParams)
}
