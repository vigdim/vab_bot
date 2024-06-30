package methods

import (
	"github.com/vigdim/vab_bot/database"
	"strings"
)

// GetUser - Получаем пользователя
func GetUser(TgId int64) ([]models.Users, bool) {
	var User []models.Users
	result := models.DB.First(&User, "tg_id = ?", TgId) // Ищем в базе Users запись с TgId текущего пользователя
	UserYesNo := result.RowsAffected > 0
	return User, UserYesNo
}

// GetDbAllOfd - Получаем всех ОФД операторов и их количество
func GetDbAllOfd() ([]models.Ofd, int64) {
	var Ofds []models.Ofd
	result := models.DB.Find(&Ofds)
	return Ofds, result.RowsAffected
}

// GetDbOneOfd - Получаем всех ОФД операторов
func GetDbOneOfd(data_query string) ([]models.Code, string, string) {

	ofd_name := strings.Split(data_query, "_")[2]
	ofd_id := strings.Split(data_query, "_")[4]

	var Code []models.Code
	models.DB.Distinct("period_id", "price_id").Where("ofd_id = ?", ofd_id).Order("price_id asc").Find(&Code)

	return Code, ofd_name, ofd_id
}

func GetValuesCodesToString(OfdId string, PeriodId string, PriceId string) (string, string, string) {
	//var Code []models.Code
	var Ofd []models.Ofd
	var Period []models.Period
	var Price []models.Price

	_ = models.DB.First(&Ofd, "id = ?", OfdId)
	_ = models.DB.First(&Period, "id = ?", PeriodId)
	_ = models.DB.First(&Price, "id = ?", PriceId)

	return Ofd[0].OfdName, Period[0].PeriodName, Price[0].Price
}

// GetDbBuyOfd - Получаем все купленные коды ОФД
func GetDbBuyOfd(TgId int64) ([]models.Buy, int64) {
	var Buy []models.Buy
	result := models.DB.Where("tg_id = ?", TgId).Order("id asc").Find(&Buy)

	return Buy, result.RowsAffected
}

func GetSetPayStatus(TgId string, PayID string, PayNotify string, nameOfd string, OfdId string, namePeriod string,
	PeriodId string, namePrice string, PriceId string) error {

	var PayStatus models.PayStatus
	result := models.DB.First(&PayStatus, "pay_id = ?", PayID)
	if result.RowsAffected == 0 { // Если не находим, то создаем его
		PayStatus = models.PayStatus{TgId: TgId, PayID: PayID, PayNotify: PayNotify, OfdName: nameOfd, OfdID: OfdId,
			PeriodID: PeriodId, PeriodName: namePeriod, PriceName: namePrice, PriceID: PriceId}
		result = models.DB.Create(&PayStatus)
	} else { // Иначе обновляем данные
		// Изменения действуют только на ненулевые приходные переменные (пустые переменные не трогают имеющиеся записи в строке!
		result = models.DB.Model(&PayStatus).Updates(models.PayStatus{TgId: TgId, PayID: PayID, PayNotify: PayNotify,
			OfdName: nameOfd, OfdID: OfdId, PeriodID: PeriodId, PeriodName: namePeriod, PriceName: namePrice, PriceID: PriceId})
	}

	return result.Error
}
