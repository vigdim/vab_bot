package user_handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"vab/keyboards"
	"vab/utils"
)

// UsHi Вывод приветствия
//func UsHi(bot *telego.Bot, update telego.Update) {
//	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
//		fmt.Sprintf("Привет! %s", methods.GetUser())).WithReplyMarkup(keyboards.Kb_user_Hi))
//}

func CatalogMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), "<b>Меню - 📁Каталог</b>").WithReplyMarkup(keyboards.Kb_catalog).WithParseMode(telego.ModeHTML))
}

func MainMenuMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), "<b>Главное 🍽 меню</b>").WithReplyMarkup(keyboards.Kb_main).WithParseMode(telego.ModeHTML))
}
