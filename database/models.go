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
	//err = DB.AutoMigrate(&Ofd{}, &Period{}, &Price{}, &Buy{}, &Code{}, &Users{})
	//if err != nil {log.Fatalf("Ошибка AutoMigrate функции %s", err)}
}

// БЛОК ОПИСАНИЯ ТАБЛИЦ ///////////////////////////////////////////////////////////////////////////
// Имена таблиц и колонок в реальности создаются в нижнем регистре ================================

// Users Таблица users
type Users struct {
	gorm.Model
	Name     string
	TgID     string
	Role     string
	Email    string
	Discount int
}

// Ofd Таблица ofds
type Ofd struct {
	gorm.Model
	OfdName string
	OfdLink string
	Code    Code
}

// Period Таблица periods
type Period struct {
	gorm.Model
	PeriodName string
	Code       Code
}

// Price Таблица prices
type Price struct {
	gorm.Model
	Price string
	Code  Code
}

// Buy Таблица buys
type Buy struct {
	gorm.Model
	TgId       string
	Email      string
	OfdName    string
	CodeNum    string
	PeriodName string
	Price      uint64
}

// Code Таблица codes
type Code struct {
	gorm.Model
	CodeNum  string
	OfdID    uint
	PeriodID uint
	PriceID  uint
}

// КОНЕЦ БЛОКА ОПИСАНИЯ ТАБЛИЦ ////////////////////////////////////////////////////////////////////
