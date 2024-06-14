package user_handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"vab/database/methods"
	"vab/keyboards"
	"vab/utils"
)

// UsHi Вывод приветствия
//func UsHi(bot *telego.Bot, update telego.Update) {
//	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
//		fmt.Sprintf("Привет! %s", methods.GetUser())).WithReplyMarkup(keyboards.Kb_user_Hi))
//}

// CatalogMess Меню - Каталог
func CatalogMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Меню - 📁Каталог</b>").WithReplyMarkup(keyboards.Kb_catalog).WithParseMode(telego.ModeHTML))
}

// MainMenuMess - Главное меню
func MainMenuMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Главное 🍽 меню</b>").WithReplyMarkup(keyboards.Kb_main).WithParseMode(telego.ModeHTML))
}

// PayOfd - Меню покупки ОФД
func PayOfd(bot *telego.Bot, update telego.Update) {

	var ofds, len_ofds = methods.GetAllOfd()

	var bnOfd = make([]telego.InlineKeyboardButton, len_ofds)
	var bnDesc = make([]telego.InlineKeyboardButton, len_ofds)

	row1 := tu.InlineKeyboardRow()
	row2 := tu.InlineKeyboardRow()
	row3 := tu.InlineKeyboardRow()
	row4 := tu.InlineKeyboardRow()
	row5 := tu.InlineKeyboardRow()
	row6 := tu.InlineKeyboardRow()

	for index, value := range ofds {
		//fmt.Println(index, value)
		bnOfd[index] = tu.InlineKeyboardButton(value.OfdName).WithCallbackData("cback_" + value.OfdName)
		bnDesc[index] = tu.InlineKeyboardButton("Описание").WithWebApp(tu.WebAppInfo(value.OfdLink))
	}
	row1 = append(row1, bnOfd[0], bnDesc[0])
	row2 = append(row2, bnOfd[1], bnDesc[1])
	row3 = append(row3, bnOfd[2], bnDesc[2])
	row4 = append(row4, bnOfd[3], bnDesc[3])
	row5 = append(row5, bnOfd[4], bnDesc[4])
	row6 = append(row6, bnOfd[5], bnDesc[5])

	inlineKeyboard := tu.InlineKeyboard(row1, row2, row3, row4, row5, row6)

	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Выберите оператора ОФД</b>").WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
}
