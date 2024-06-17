package methods

import (
	"vab/database"
)

// AddUser - Создаем пользователя
func AddUser(TgId string, Email string, Name string, Discount int, Role string) error {
	var user models.Users
	result := models.DB.First(&user.TgID, TgId) // Ищем в базе Users запись с TgId текущего пользователя

	if result.RowsAffected == 0 { // Если не находим, то создаем его
		user := models.Users{TgID: TgId, Email: Email, Name: Name, Discount: Discount, Role: Role}
		result = models.DB.Create(&user)
	}
	return result.Error
}
