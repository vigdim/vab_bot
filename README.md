main.go -> func main()

bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

При продакшене удалить/закомментить функцию - telego.WithDefaultDebugLogger()
