package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB // База данных с инициализацией в func Init()

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	SQLCON := os.Getenv("MYSQLCON")

	DB, err = gorm.Open(mysql.Open(SQLCON), &gorm.Config{})
	//err = DB.AutoMigrate(&Ofd{}, &Period{}, &Price{}, &Buy{}, &Code{})
	if err != nil {
		log.Fatalf("Ошибка AutoMigrate функции %s", err)
	}
	//DB.Create(&Users{Name: "Dmitry"})

	//var result Ofd
	//DB.Where("ofd_name = ?", "One").First(&result)
	//DB.Create(&Code{
	//	Model:    gorm.Model{},
	//	CodeNum:  "156723",
	//	OfdID:    result.ID,
	//	PeriodID: 1,
	//	PriceID:  1,
	//})

	var resultCode Code
	DB.Where("code_num = ?", "156723").First(&resultCode)
	var resultOfd Ofd
	DB.Where("id = ?", resultCode.OfdID).First(&resultOfd)
	var resultPeriod Period
	DB.Where("id = ?", resultCode.PeriodID).First(&resultPeriod)
	var resultPrice Price
	DB.Where("id = ?", resultCode.PeriodID).First(&resultPrice)
	price, _ := strconv.ParseFloat(resultPrice.Price, 32)

	tgId := "1234567890"
	DB.Create(&Buy{
		Model:      gorm.Model{},
		TgId:       tgId,
		Email:      "aaa@bbb.com",
		OfdName:    resultOfd.OfdName,
		CodeNum:    resultCode.CodeNum,
		PeriodName: resultPeriod.PeriodName,
		Price:      float32(price),
	})

	//=============
	//DB.Where("code_num = ?", "333333555555").Delete(&Code{}) // Помечено как удаленное с возможностью восстановления
	DB.Where("code_num = ?", "333333555555").Unscoped().Delete(&Code{})
}

// БЛОК ОПИСАНИЯ ТАБЛИЦ ///////////////////////////////////////////////////////////////////////////
// Имена таблиц и колонок в реальности создаются в нижнем регистре ================================
type Users struct { // Таблица users
	gorm.Model
	Name string
	TgID string
	Role string
}

type Ofd struct { // Таблица ofds
	gorm.Model
	OfdName string
	OfdLink string
	Code    Code
}

type Period struct { // Таблица periods
	gorm.Model
	PeriodName string
	Code       Code
}

type Price struct { // Таблица prices
	gorm.Model
	Price string
	Code  Code
}

type Buy struct { // Таблица buys
	gorm.Model
	TgId       string
	Email      string
	OfdName    string
	CodeNum    string
	PeriodName string
	Price      float32
}

type Code struct { // Таблица codes
	gorm.Model
	CodeNum  string
	OfdID    uint
	PeriodID uint
	PriceID  uint
}

// КОНЕЦ БЛОКА ОПИСАНИЯ ТАБЛИЦ ////////////////////////////////////////////////////////////////////
