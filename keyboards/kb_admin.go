package keyboards

import tu "github.com/mymmrac/telego/telegoutil"

// Kb_admin_first - Начальное меню админа
var Kb_admin_first = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("🥷️Админ"),
		tu.KeyboardButton("👨‍💼Пользователь"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")
