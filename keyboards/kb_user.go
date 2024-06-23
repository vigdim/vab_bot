package keyboards

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

// Kb_main - Главное меню юзера
var Kb_main = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("📁 Каталог"),
		tu.KeyboardButton("🔐 Кабинет"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")

// Kb_catalog - Меню каталога сервисов
var Kb_catalog = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("💵 Купить ОФД"),
		tu.KeyboardButton("💬 Консультация"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("©️О компании"),
		tu.KeyboardButton("🔙 В НАЧАЛО"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")

// Kb_сabinet - Меню кабинета
var Kb_сabinet = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("👤 Аккаунт"),
		tu.KeyboardButton("🧾 Заказы"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("🔙 В НАЧАЛО"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")
