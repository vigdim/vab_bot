package main

import (
	"github.com/joho/godotenv" // Загрузка настроек из .env файла
	models "github.com/vigdim/vab_bot/database"
	botrouter "github.com/vigdim/vab_bot/router/bot"
	siterouter "github.com/vigdim/vab_bot/router/http"
	"github.com/vigdim/vab_bot/utils"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	utils.DOMAIN_SERVER = os.Getenv("DOMAIN_SERVER")
	utils.ACQUIRING_SERVER = os.Getenv("ACQUIRING_SERVER")
	utils.CHECK_SERVER = os.Getenv("CHECK_SERVER")

	models.Init()
	go botrouter.Init()
	siterouter.Init()
}
