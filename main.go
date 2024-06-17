package main

import (
	"fmt"
	"github.com/joho/godotenv" // Загрузка настроек из .env файла
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
	"os"
	"vab/database"
	"vab/handlers"
	admin_handlers "vab/handlers/admin"
	"vab/handlers/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	botToken := os.Getenv("TOKEN")

	models.Init()

	// Примечание. Средство ведения журнала по умолчанию может раскрывать конфиденциальную информацию.
	// Использовать только в разработке
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
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
	// Переход в меню Консультация
	bh.Handle(user_handlers.ConsultationMess, th.TextEqual("💬 Консультация"))
	// Переход в меню Купить ОФД (вывод списка ОФД)
	bh.Handle(user_handlers.ListOfd, th.TextEqual("💵 Купить ОФД"))

	// user handlers CallbackQuery =======================================================
	// cback_ОФД
	bh.HandleCallbackQuery(user_handlers.GetOneOfdCb, th.AnyCallbackQueryWithMessage(), th.CallbackDataPrefix("cb_OFD_"))
	// cback Консультация
	bh.HandleCallbackQuery(user_handlers.ConsultationCb, th.AnyCallbackQueryWithMessage(), th.CallbackDataPrefix("cb_cons_"))

	// admin handlers message ==============================================================
	// Переход в меню Каталог
	bh.Handle(admin_handlers.MainUserMenuMess, th.TextEqual("👨‍💼Пользователь"))

	// ============================================================================================
	// main_handlers ==============================================================================
	bh.Handle(main_handlers.Start, th.CommandEqual("start"))    //Вывод приветствия
	bh.Handle(main_handlers.SendMyData, th.TextEqual("telega")) //Вывод данных пользователя
	//bh.Handle(handlers.SendLogo, th.TextEqual("88"))    //Вывод логотипа
	// END HANDLERS BLOCK /////////////////////////////////////////////////////////////////////////

	// TODO: Clean code
	//http.HandleFunc("/", sayhello)       // Устанавливаем роутер
	//err = http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	//http.HandleFunc("/", sayhello)
	//go http.ListenAndServe(":80", nil)

	// Start handling updates
	bh.Start()
}
