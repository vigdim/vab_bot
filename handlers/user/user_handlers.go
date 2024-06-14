package user_handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"strconv"
	"strings"
	models "vab/database"
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

	var ofds, len_ofds = methods.GetDbAllOfd()

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
		bnOfd[index] = tu.InlineKeyboardButton(value.OfdName).WithCallbackData("cb_OFD_" + value.OfdName + "_id_" + strconv.Itoa(int(value.ID)))
		bnDesc[index] = tu.InlineKeyboardButton("Описание").WithWebApp(tu.WebAppInfo(value.OfdLink))
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("What about C#?"),

			tu.InlineKeyboardButton("LOL").WithCallbackData("LOL"),
		)
	}

	row1 = append(row1, bnOfd[0], bnDesc[0])
	row2 = append(row2, bnOfd[1], bnDesc[1])
	row3 = append(row3, bnOfd[2], bnDesc[2])
	row4 = append(row4, bnOfd[3], bnDesc[3])
	row5 = append(row5, bnOfd[4], bnDesc[4])
	row6 = append(row6, bnOfd[5], bnDesc[5])

	//type Rows struct {
	//	row []telego.InlineKeyboardButton
	//}

	inlineKeyboard := tu.InlineKeyboard(row1, row2, row3, row4, row5, row6)

	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Выберите оператора ОФД</b>").WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
}

// MainMenuMess - Главное меню
func GetOneOfd(bot *telego.Bot, query telego.CallbackQuery) {
	var Code, ofd_name = methods.GetDbOneOfd(query.Data)

	_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText(strings.Split(query.Data, "_")[2]))

	var PeriodName []models.Period
	var Price []models.Price

	for index, _ := range Code {
		models.DB.First(&PeriodName, Code[index].PeriodID)
		models.DB.First(&Price, Code[index].PriceID)
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			"Код активации "+ofd_name+" ОФД\n<b>Срок: "+PeriodName[0].PeriodName+"</b>"+
				"\n<b>Цена: "+Price[0].Price+" рублей</b>").WithParseMode(telego.ModeHTML))

		//_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), PeriodName[0].PeriodName))
	}
}
