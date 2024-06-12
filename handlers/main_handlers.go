package main_handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"vab/keyboards"
)

// Start Register new handler with match on command `/start`
func Start(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Привет %s!\nДобро пожаловать в сервис ККМ от компании ВАБ!",
			update.Message.From.FirstName)).WithReplyMarkup(keyboards.Kb_main))
}

// Register new handler with match 88
//file := utils.MustOpen("files/vab.png")
//func SendLogo(bot *telego.Bot, update telego.Update) {
//	_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
//}
