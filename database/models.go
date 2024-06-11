package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB // База данных с инициализацией в func Init()

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	SQLCON := os.Getenv("MYSQLCON")

	DB, err = gorm.Open(mysql.Open(SQLCON), &gorm.Config{})
	err = DB.AutoMigrate(&Users{})
	if err != nil {
		log.Fatalf("Ошибка AutoMigrate функции %s", err)
	}
	//DB.Create(&Users{Name: "Dmitry"})
}

// БЛОК ОПИСАНИЯ ТАБЛИЦ ///////////////////////////////////////////////////////////////////////////
// Имена таблиц и колонок в реальности создаются в нижнем регистре ================================
type Users struct { // Таблица users
	gorm.Model
	Name string
	TgID string
	Role string
}

// КОНЕЦ БЛОКА ОПИСАНИЯ ТАБЛИЦ ////////////////////////////////////////////////////////////////////
