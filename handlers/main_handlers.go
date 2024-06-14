package main_handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"vab/keyboards"
	"vab/utils"
)

// Start Register new handler with match on command `/start`
func Start(bot *telego.Bot, update telego.Update) {
	file := utils.MustOpen("files/vab.png")
	_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("<b>Здравствуйте, %s!\nДобро пожаловать в сервис ККТ от компании ВАБ!</b>",
			update.Message.From.FirstName)).WithReplyMarkup(keyboards.Kb_main).WithParseMode(telego.ModeHTML))
}

// Register new handler with match 88
//file := utils.MustOpen("files/vab.png")
//func SendLogo(bot *telego.Bot, update telego.Update) {
//	_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
//}
