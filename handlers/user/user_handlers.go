package user_handlers

import (
	"fmt"
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

	//inlineKeyboard := tu.InlineKeyboard(
	//	tu.InlineKeyboardRow( // Row 1
	//		tu.InlineKeyboardButton("methods.GetUser()"). // Column 1
	//								WithCallbackData("callback_1"),
	//		tu.InlineKeyboardButton(ofds[5].OfdName). // Column 2
	//								WithCallbackData("callback_2"),
	//	),
	//	tu.InlineKeyboardRow( // Row 2
	//		tu.InlineKeyboardButton("URL button").WithURL("https://example.com"), // Column 1
	//	),
	//)
	//
	//

	var ofds, len_ofds = methods.GetAllOfd()

	var bnOfd = make([]telego.InlineKeyboardButton, len_ofds)
	var bnDesc = make([]telego.InlineKeyboardButton, len_ofds)

	//var num_rows = len_ofds / 2 + len_ofds % 2
	//var rows = make([]telego.InlineKeyboardButton, len_ofds)

	for index, value := range ofds {
		fmt.Println(index, value)
		bnOfd[index] = tu.InlineKeyboardButton(value.OfdName).WithCallbackData("cback_" + value.OfdName)
		bnDesc[index] = tu.InlineKeyboardButton("Описание").WithWebApp(tu.WebAppInfo(value.OfdLink))

		//rows[index] = tu.InlineKeyboardButton(value.OfdName).WithCallbackData("cback_" + value.OfdName)

	}

	row1 := tu.InlineKeyboardRow()
	//row1 = tu.InlineKeyboardRow()

	//row2 := tu.InlineKeyboardRow()
	//
	//rows = append(row1, bnOfd[0], bnDesc[0])
	//row2 = append(row2, bnOfd[1], bnDesc[1])

	inlineKeyboard := tu.InlineKeyboard()

	for i := 0; i < 3; i++ {
		row1 = append(tu.InlineKeyboardRow(), bnOfd[i], bnDesc[i])
		inlineKeyboard = tu.InlineKeyboard(row1)
	}
	users := []string{"Tom", "Alice", "Kate"}
	users = append(users, "Bob")
	//utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Выберите оператора ОФД</b>").WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
}
