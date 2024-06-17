package main_handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"strconv"
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

// Start Register new handler with match on command `/start`
func SendMyData(bot *telego.Bot, update telego.Update) {
	tg_id := update.Message.From.ID
	fname := update.Message.From.FirstName
	lname := update.Message.From.LastName
	uname := update.Message.From.Username
	lang := update.Message.From.LanguageCode
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("<b>Твои данные:</b>\nTelegram id: <b><code>%s</code></b>\nFirst name: <b>%s</b>"+
			"\nLast name: <b>%s</b>\nUser name: <b>%s</b>\nLanguage code: <b>%s</b>",
			strconv.Itoa(int(tg_id)), fname, lname, uname, lang)).WithParseMode(telego.ModeHTML))
}
