package user_servises

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"vab/database/methods"
	"vab/keyboards"
)

// UsHi Вывод приветствия
func UsHi(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Привет! %s", methods.GetUser())).WithReplyMarkup(keyboards.Kb_user_Hi))
}
