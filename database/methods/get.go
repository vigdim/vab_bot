package methods

import (
	models "vab/database"
)

// GetUser Получаем пользователя
func GetUser() string {
	var result models.Users
	models.DB.First(&result, 1)
	return result.Name
}
