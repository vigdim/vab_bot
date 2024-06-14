package methods

import (
	"strings"
	models "vab/database"
)

// GetUser - Получаем пользователя
func GetUser() string {
	var result models.Users
	models.DB.First(&result, 1)
	return result.Name
}

// GetDbAllOfd - Получаем всех ОФД операторов и их количество
func GetDbAllOfd() ([]models.Ofd, int64) {
	var Ofds []models.Ofd
	result := models.DB.Find(&Ofds)
	return Ofds, result.RowsAffected
}

// GetDbOneOfd - Получаем всех ОФД операторов и их количество
func GetDbOneOfd(data_query string) ([]models.Code, string) {

	ofd_name := strings.Split(data_query, "_")[2]
	ofd_id := strings.Split(data_query, "_")[4]

	var Code []models.Code
	models.DB.Distinct("period_id", "price_id").Where("ofd_id = ?", ofd_id).Order("price_id asc").Find(&Code)

	//models.DB.Where("ofd_id = ?", ofd_id).First(&Code)
	return Code, ofd_name
}
