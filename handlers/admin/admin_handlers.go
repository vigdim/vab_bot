package admin_handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"vab/keyboards"
	"vab/utils"
)

// MainUserMenuMess - Главное меню пользователя
func MainUserMenuMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Главное 🍽 меню</b>").WithReplyMarkup(keyboards.Kb_main).WithParseMode(telego.ModeHTML))
}
