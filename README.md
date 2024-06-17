# About2
Бот по продаже лицензий ОФД и Каталог товаров и услуг

main.go -> func main()

bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

При продакшене удалить/закомментить функцию - telego.WithDefaultDebugLogger()

