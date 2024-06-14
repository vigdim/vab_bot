package main

import (
	"fmt"
	"github.com/joho/godotenv" // Загрузка настроек из .env файла
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	models "vab/database"
	"vab/handlers"
	user_handlers "vab/handlers/user"

	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	botToken := os.Getenv("TOKEN")

	models.Init()
	//SQLCON := os.Getenv("MYSQLCON")
	//db, err := gorm.Open(mysql.Open(SQLCON), &gorm.Config{})
	//db.AutoMigrate(&models.User4{})
	//var result User4
	//db.First(&result, 1)
	//db.Create(&User4{Name: "Dmitry"})

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
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
	// user_services handlers message ==============================================================
	// Переход в меню Каталог
	bh.Handle(user_handlers.CatalogMess, th.TextEqual("📁 Каталог"))
	// Переход в меню Каталог
	bh.Handle(user_handlers.MainMenuMess, th.TextEqual("🔙 В НАЧАЛО"))
	// Переход в меню Купить ОФД
	bh.Handle(user_handlers.PayOfd, th.TextEqual("💵 Купить ОФД"))

	// user_services handlers CallbackQuery =======================================================
	// cback_ОФД
	bh.HandleCallbackQuery(user_handlers.GetOneOfd, th.AnyCallbackQueryWithMessage(), th.CallbackDataPrefix("cb_OFD_"))
	// ============================================================================================

	// main_handlers ==============================================================================
	bh.Handle(main_handlers.Start, th.CommandEqual("start")) //Вывод приветствия
	//bh.Handle(handlers.SendLogo, th.TextEqual("88"))    //Вывод логотипа
	// END HANDLERS BLOCK /////////////////////////////////////////////////////////////////////////

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		fmt.Println("Start")
	}, th.TextPrefix("st"))

	// Register new handler with match on command `/sticker`
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendSticker(tu.Sticker(tu.ID(update.Message.Chat.ID), tu.FileFromID("CAACAgIAAxkBAAEoaBZlhvb7_eVfDNHhdvVb9zOS834EYAACgQEAAiteUwteCmw-bAABeLQzBA")))
	}, th.CommandEqual("sticker"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(tu.Message(tu.ID(update.Message.Chat.ID), "Unknown command, use /start"))
	}, th.TextSuffix("90"))

	//file := utils.MustOpen("files/vab.png")
	//bh.Handle(func(bot *telego.Bot, update telego.Update) {
	//	_, _ = bot.SendPhoto(tu.Photo(tu.ID(update.Message.Chat.ID), tu.File(file)))
	//}, th.TextSuffix("77"))

	// Start handling updates
	bh.Start()
}

// Helper function to open file or panic
//func mustOpen(filename string) *os.File {
//	file, err := os.Open(filename)
//	if err != nil {
//		panic(err)
//	}
//	return file
//}
