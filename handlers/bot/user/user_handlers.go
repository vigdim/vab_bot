package user_handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vigdim/vab_bot/database"
	"github.com/vigdim/vab_bot/database/methods"
	"github.com/vigdim/vab_bot/keyboards"
	"github.com/vigdim/vab_bot/utils"
	"os"
	"strconv"
	"strings"
)

// CatalogMess Меню - Каталог
func CatalogMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Меню 📁 Каталог</b>").WithReplyMarkup(keyboards.Kb_catalog).WithParseMode(telego.ModeHTML))
}

// AboutMess Меню - О компании
func AboutMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	file := utils.MustOpen("files/vab.png")
	_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		utils.AboutCompany).
		WithReplyMarkup(keyboards.Kb_catalog).WithParseMode(telego.ModeHTML))
}

// MainMenuMess - Главное меню
func MainMenuMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Главное 🍽 меню</b>").WithReplyMarkup(keyboards.Kb_main).WithParseMode(telego.ModeHTML))
}

// CabinetMess - Меню Кабинет
func CabinetMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Меню 🔐 Кабинет</b>").WithReplyMarkup(keyboards.Kb_сabinet).WithParseMode(telego.ModeHTML))
}

// AccountMess Меню - Аккаунт
func AccountMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Меню 👤 Аккаунт</b>").WithReplyMarkup(keyboards.Kb_сabinet).WithParseMode(telego.ModeHTML))
	strUserId := strconv.Itoa(int(update.Message.Chat.ID))
	CurrentUser, UserYesNo := methods.GetUser(update.Message.Chat.ID)
	// Если пользователя не существует в базе
	if UserYesNo == false {
		err := methods.AddUser(strUserId, "", "", "", 0, "user")
		if err != nil {
			return
		}
		CurrentUser, UserYesNo = methods.GetUser(update.Message.Chat.ID)
	}
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("📰 Внести свои данные").WithWebApp(tu.WebAppInfo(utils.DOMAIN + "/account?us_id=" +
				strUserId + "&name=" + CurrentUser[0].Name + "&email=" + CurrentUser[0].Email + "&phone=" + CurrentUser[0].Phone)),
		),
	)
	mess, _ := bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Внесите свои данные для получения чеков покупки на почту и для возможной обратной связи:</b>").WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
	// Помещаем сообщение на удаление с чата (нужно удалять так как кнопка будет содержать устаревшие данные
	utils.MessIdForDel = mess.MessageID
	utils.CurUpdate = update      // Инициализируем для вызова из функции сайта site_router->AccountPost
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
}

// ListOfd - Меню покупки ОФД
func ListOfd(bot *telego.Bot, update telego.Update) {
	var (
		ofds, lenOfds = methods.GetDbAllOfd()
		bnOfd         = make([]telego.InlineKeyboardButton, lenOfds)
		bnDesc        = make([]telego.InlineKeyboardButton, lenOfds)
		grid          = make([][]telego.InlineKeyboardButton, lenOfds)
	)

	for index, value := range ofds {
		bnOfd[index] = tu.InlineKeyboardButton(value.OfdName).WithCallbackData("cb_OFD_" + value.OfdName + "_id_" + strconv.Itoa(int(value.ID)))
		bnDesc[index] = tu.InlineKeyboardButton("Описание").WithWebApp(tu.WebAppInfo(value.OfdLink))
		grid[index] = append(grid[index], bnOfd[index], bnDesc[index])
	}

	inlineKeyboard := tu.InlineKeyboardGrid(grid)
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Выберите оператора ОФД:</b>").WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
}

// GetOneOfdCb - щелчек по кнопке оператора ОФД
func GetOneOfdCb(bot *telego.Bot, query telego.CallbackQuery) {
	var (
		PeriodName     []models.Period
		Price          []models.Price
		inlineKeyboard = tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("💵 Оплатить код активации 💵").
					WithCallbackData("callback_payofd"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("🧾 Получить счёт на оплату 🧾").
					WithCallbackData("callback_orderofd"),
			),
		)
	)

	var Code, ofd_name = methods.GetDbOneOfd(query.Data)
	if len(Code) == 0 {
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			"️❗️Коды активации для оператора "+strings.Split(query.Data, "_")[2]+
				" ОФД временно недоступны.\nВойдите в меню 💬 <b>Консультация</b> для решения проблемы.").WithParseMode(telego.ModeHTML))
	}
	_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText(strings.Split(query.Data, "_")[2]))

	for index := range Code {
		models.DB.First(&PeriodName, Code[index].PeriodID)
		models.DB.First(&Price, Code[index].PriceID)
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			"Код активации "+ofd_name+" ОФД\n<b>Срок: "+PeriodName[0].PeriodName+"</b>"+
				"\n<b>Цена: "+Price[0].Price+" рублей</b>").WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
	}
}

// ConsultationMess - Меню Консультация
func ConsultationMess(bot *telego.Bot, update telego.Update) {
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Приобретение ККТ").WithCallbackData("cb_cons_PayKKT"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Приобретение ФН или ОФД").WithCallbackData("cb_cons_PayFnOfd"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Обслуживание ККТ").WithCallbackData("cb_cons_ServKKT"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Бухгалтерское обслуживание").WithCallbackData("cb_cons_ServBuh"),
		),
	)

	Phone := os.Getenv("VAB_PHONE")
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		"<b>Меню 💬 Консультация</b>").WithParseMode(telego.ModeHTML))
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(utils.Consultation,
			Phone)).WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))

}

// ConsultationCb - Обработка кнопок консультации
func ConsultationCb(bot *telego.Bot, query telego.CallbackQuery) {
	var choice = strings.Split(query.Data, "_")[2]
	VabTgId, _ := strconv.ParseInt(os.Getenv("VAB_TG_ID"), 10, 64)
	inlineKeyboardPeriod := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("🔂 Ответить в личку").WithURL("https://t.me/"+query.From.Username),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("🕝 10 минут").
				WithCallbackData("cb_cons_Period_"+strconv.Itoa(int(query.From.ID))+"_10 МИНУТ"),
			tu.InlineKeyboardButton("🕝 20 минут").
				WithCallbackData("cb_cons_Period_"+strconv.Itoa(int(query.From.ID))+"_20 МИНУТ"),
			tu.InlineKeyboardButton("🕝 30 минут").
				WithCallbackData("cb_cons_Period_"+strconv.Itoa(int(query.From.ID))+"_30 МИНУТ"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("🕝 1 час").
				WithCallbackData("cb_cons_Period_"+strconv.Itoa(int(query.From.ID))+"_ЧАСА"),
			tu.InlineKeyboardButton("🕝 1 день").
				WithCallbackData("cb_cons_Period_"+strconv.Itoa(int(query.From.ID))+"_ДНЯ"),
			tu.InlineKeyboardButton("🕝 В раб. день").
				WithCallbackData("cb_cons_Period_"+strconv.Itoa(int(query.From.ID))+"_РАБОЧЕГО ДНЯ"),
		),
	)
	switch choice {
	case "PayKKT":
		inlineKeyboard := utils.ConfirmSend("cb_cons_PayKKTYes")
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationPayKKT).WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
	case "PayKKTYes":
		_, _ = bot.SendSticker(tu.Sticker(tu.ID(VabTgId),
			tu.FileFromID("CAACAgIAAxkBAAEoaBZlhvb7_eVfDNHhdvVb9zOS834EYAACgQEAAiteUwteCmw-bAABeLQzBA")))
		_, _ = bot.SendMessage(tu.Message(tu.ID(VabTgId),
			fmt.Sprintf(utils.ConsultationPayKKTYes, query.From.FirstName, query.From.LastName)).
			WithReplyMarkup(inlineKeyboardPeriod).WithParseMode(telego.ModeHTML))
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationMesSend).
			WithParseMode(telego.ModeHTML))
	case "PayFnOfd":
		inlineKeyboard := utils.ConfirmSend("cb_cons_PayFnOfdYes")
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationPayFnOfd).WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
	case "PayFnOfdYes":
		_, _ = bot.SendSticker(tu.Sticker(tu.ID(VabTgId),
			tu.FileFromID("CAACAgIAAxkBAAEoaBZlhvb7_eVfDNHhdvVb9zOS834EYAACgQEAAiteUwteCmw-bAABeLQzBA")))
		_, _ = bot.SendMessage(tu.Message(tu.ID(VabTgId),
			fmt.Sprintf(utils.ConsultationPayFnOfdYes, query.From.FirstName, query.From.LastName)).
			WithReplyMarkup(inlineKeyboardPeriod).WithParseMode(telego.ModeHTML))
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationMesSend).
			WithParseMode(telego.ModeHTML))
	case "ServKKT":
		inlineKeyboard := utils.ConfirmSend("cb_cons_ServKKTYes")

		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationServKKT).WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
	case "ServKKTYes":
		_, _ = bot.SendSticker(tu.Sticker(tu.ID(VabTgId),
			tu.FileFromID("CAACAgIAAxkBAAEoaBZlhvb7_eVfDNHhdvVb9zOS834EYAACgQEAAiteUwteCmw-bAABeLQzBA")))
		_, _ = bot.SendMessage(tu.Message(tu.ID(VabTgId),
			fmt.Sprintf(utils.ConsultationServKKTYes, query.From.FirstName, query.From.LastName)).
			WithReplyMarkup(inlineKeyboardPeriod).WithParseMode(telego.ModeHTML))
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationMesSend).
			WithParseMode(telego.ModeHTML))
	case "ServBuh":
		inlineKeyboard := utils.ConfirmSend("cb_cons_ServBuhYes")

		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationServBuh).WithReplyMarkup(inlineKeyboard).WithParseMode(telego.ModeHTML))
	case "ServBuhYes":
		_, _ = bot.SendSticker(tu.Sticker(tu.ID(VabTgId),
			tu.FileFromID("CAACAgIAAxkBAAEoaBZlhvb7_eVfDNHhdvVb9zOS834EYAACgQEAAiteUwteCmw-bAABeLQzBA")))
		_, _ = bot.SendMessage(tu.Message(tu.ID(VabTgId),
			fmt.Sprintf(utils.ConsultationServBuhYes, query.From.FirstName, query.From.LastName)).
			WithReplyMarkup(inlineKeyboardPeriod).WithParseMode(telego.ModeHTML))
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID),
			utils.ConsultationMesSend).
			WithParseMode(telego.ModeHTML))
	case "Period":
		sds := strings.Split(query.Data, "_")[3]
		ids, _ := strconv.ParseInt(sds, 10, 64)
		answ_period := strings.Split(query.Data, "_")[4]
		_, _ = bot.SendMessage(tu.Message(tu.ID(VabTgId),
			fmt.Sprintf("👍 <b>Отправлено:</b> связь в течении %s", answ_period)).WithParseMode(telego.ModeHTML))

		AnswerConsultation(bot, ids, answ_period)
	}
}

func AnswerConsultation(bot *telego.Bot, id_ans int64, time_ans string) {
	VabPhone := os.Getenv("VAB_PHONE")
	_, _ = bot.SendSticker(tu.Sticker(tu.ID(id_ans),
		tu.FileFromID("CAACAgIAAxkBAAEsBoFmbp43OPkseIFtScovWoy2ehKd1wACjwEAAiteUwtSWZbMUSUiwjUE")))
	_, _ = bot.SendMessage(tu.Message(tu.ID(id_ans),
		fmt.Sprintf("💬 <b>Ваше сообщение получено!</b>\n<i>В течении</i> <b>%s</b> <i>мы свяжемся с Вами через личное"+
			" сообщение или позвоним, если Ваш номер не скрыт в Telegram. 😊\n"+
			"Также Вы можете самостоятельно позвонить нам по номеру:</i> %s", time_ans, VabPhone)).WithParseMode(telego.ModeHTML))
}

// OrdersMess - Переход в меню Заказы (вывод списка купленных ОФД кодов)
func OrdersMess(bot *telego.Bot, update telego.Update) {
	var Buys, lenBuys = methods.GetDbBuyOfd(update.Message.From.ID)
	utils.DelMessage(bot, update) // Удаляем предыдущее сообщение
	_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("<b>💵 Приобретенные ранее коды ОФД:</b>")).WithParseMode(telego.ModeHTML))
	if lenBuys > 0 {
		for index, value := range Buys {
			_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("%s. <b>%s</b> Дата: %s\nПериод: %s. Код: <code>%s</code>\n<b><a href=\"%s\">Смотреть чек оплаты</a></b>",
					strconv.Itoa(index+1), value.OfdName, value.CreatedAt.Format("02-01-2006"), value.PeriodName, value.CodeNum, value.PurchaseLink)).WithParseMode(telego.ModeHTML))
		}
	} else {
		_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Приобретенные коды ОФД отсутствуют 😢.")).WithParseMode(telego.ModeHTML))
	}
}

//.WithWebApp(tu.WebAppInfo(value.OfdLink))
