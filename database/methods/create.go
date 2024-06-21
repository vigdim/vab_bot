package methods

import (
	"github.com/vigdim/vab_bot/database"
)

// AddUser - Создаем пользователя
func AddUser(TgId string, Email string, Name string, Phone string, Discount int, Role string) error {
	var user models.Users
	result := models.DB.First(&user, "tg_id = ?", TgId) // Ищем в базе Users запись с TgId текущего пользователя

	if result.RowsAffected == 0 { // Если не находим, то создаем его
		user = models.Users{TgID: TgId, Email: Email, Name: Name, Phone: Phone, Discount: Discount, Role: Role}
		result = models.DB.Create(&user)
	} else { // Иначе обновляем данные
		//user = models.Users{TgID: TgId, Email: Email, Name: Name, Phone: Phone, Discount: Discount, Role: Role}
		result = models.DB.Model(&user).Updates(models.Users{TgID: TgId, Email: Email, Name: Name, Phone: Phone, Discount: Discount, Role: Role})
	}
	return result.Error
}
