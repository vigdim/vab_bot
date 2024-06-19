package bot

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vigdim/vab_bot/database/methods"
	"github.com/vigdim/vab_bot/keyboards"
	"github.com/vigdim/vab_bot/utils"
	"strconv"
)

// Start Register new handler with match on command `/start`
func Start(bot *telego.Bot, update telego.Update) {

	var CurrentUser, UserYesNo = methods.GetUser(update.Message.From.ID)
	if UserYesNo && CurrentUser[0].Role == "admin" {
		file := utils.MustOpen("files/admin.jpg")
		_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
		_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(utils.WelcomeAdminMessage,
				update.Message.From.FirstName)).WithReplyMarkup(keyboards.Kb_admin_first).WithParseMode(telego.ModeHTML))
	} else {
		file := utils.MustOpen("files/vab.png")
		_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
		_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(utils.WelcomeUserMessage,
				update.Message.From.FirstName)).WithReplyMarkup(keyboards.Kb_main).WithParseMode(telego.ModeHTML))
	}
}

// SendMyData Вывод данных пользователя по запросу
func SendMyData(bot *telego.Bot, update telego.Update) {
	var (
		tgId         = update.Message.From.ID
		firstName    = update.Message.From.FirstName
		lastName     = update.Message.From.LastName
		username     = update.Message.From.Username
		languageCode = update.Message.From.LanguageCode
	)
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(utils.MyDataMessage, strconv.Itoa(int(tgId)), firstName, lastName, username, languageCode)).
		WithParseMode(telego.ModeHTML))
}
