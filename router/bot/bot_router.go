package bot_router

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	main_handlers "github.com/vigdim/vab_bot/handlers/bot"
	admin_handlers "github.com/vigdim/vab_bot/handlers/bot/admin"
	user_handlers "github.com/vigdim/vab_bot/handlers/bot/user"
	"os"
)

func Init() {
	botToken := os.Getenv("TOKEN")

	// Примечание. Средство ведения журнала по умолчанию может раскрывать конфиденциальную информацию.
	// Использовать только в разработке
	debugMode := os.Getenv("DEBUG_MODE")
	printErrors := os.Getenv("PRINT_ERRORS")
	var BotOption = telego.WithDefaultLogger(debugMode == "Develop", printErrors == "Yes")

	bot, err := telego.NewBot(botToken, BotOption)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// START HANDLERS BLOCK ///////////////////////////////////////////////////////////////////////
	// user handlers message ==============================================================
	// Переход в меню Каталог
	bh.Handle(user_handlers.CatalogMess, th.TextEqual("📁 Каталог"))
	// Переход в меню О компании
	bh.Handle(user_handlers.AboutMess, th.TextEqual("©️О компании"))
	// Переход в меню Каталог
	bh.Handle(user_handlers.MainMenuMess, th.TextEqual("🔙 В НАЧАЛО"))
	// Переход в меню Кабинет
	bh.Handle(user_handlers.CabinetMess, th.TextEqual("👤 Кабинет"))
	// Переход в меню Аккаунт
	bh.Handle(user_handlers.AccountMess, th.TextEqual("👤 Аккаунт"))
	// Переход в меню Консультация
	bh.Handle(user_handlers.ConsultationMess, th.TextEqual("💬 Консультация"))
	// Переход в меню Купить ОФД (вывод списка ОФД)
	bh.Handle(user_handlers.ListOfd, th.TextEqual("💵 Купить ОФД"))
	// Переход в меню Заказы (вывод списка купленных ОФД кодов)
	bh.Handle(user_handlers.OrdersMess, th.TextEqual("🧾 Заказы"))

	// user handlers CallbackQuery =======================================================
	// cback_ОФД
	bh.HandleCallbackQuery(user_handlers.GetOneOfdCb /*th.AnyCallbackQueryWithMessage(),*/, th.CallbackDataPrefix("cb_OFD_"))
	// cback Консультация
	bh.HandleCallbackQuery(user_handlers.ConsultationCb /*th.AnyCallbackQueryWithMessage(),*/, th.CallbackDataPrefix("cb_cons_"))

	// admin handlers message ==============================================================
	// Переход в меню Каталог
	bh.Handle(admin_handlers.MainUserMenuMess, th.TextEqual("👨‍💼Пользователь"))

	// main_handlers ==============================================================================
	bh.Handle(main_handlers.Start, th.CommandEqual("start"))    //Вывод приветствия
	bh.Handle(main_handlers.SendMyData, th.TextEqual("telega")) //Вывод данных пользователя

	// Start handling updates
	bh.Start()
}
