package methods

import (
	models "vab/database"
)

// GetUser - Получаем пользователя
func GetUser() string {
	var result models.Users
	models.DB.First(&result, 1)
	return result.Name
}

// GetAllOfd - Получаем всех ОФД операторов и их количество
func GetAllOfd() ([]models.Ofd, int64) {
	var Ofds []models.Ofd
	result := models.DB.Find(&Ofds)
	return Ofds, result.RowsAffected
}
