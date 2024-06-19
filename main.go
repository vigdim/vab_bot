package main

import (
	"github.com/joho/godotenv" // Загрузка настроек из .env файла
	models "github.com/vigdim/vab_bot/database"
	bot_router "github.com/vigdim/vab_bot/router/bot"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	models.Init()
	bot_router.Init()
}
