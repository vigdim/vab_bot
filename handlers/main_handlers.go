package handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// Register new handler with match on command `/start`
func Start(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Привет %s!\nДобро пожаловать в сервис ККМ от компании ВАБ!", update.Message.From.FirstName)))
}
