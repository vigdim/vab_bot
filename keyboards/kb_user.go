package keyboards

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

// Kb_main - Главное меню юзера
var Kb_main = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("📁 Каталог"),
		tu.KeyboardButton("👤 Кабинет"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")

// Kb_catalog - Меню каталога сервисов
var Kb_catalog = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("📁 Купить ОФД"),
		tu.KeyboardButton("👤 Консультация"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("🔙 В НАЧАЛО"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")

// Kb_cabinet - Меню кабинета
var Kb_сabinet = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("🛒 Корзина"),
		tu.KeyboardButton("🧾 Заказы"),
		tu.KeyboardButton("✉️ E-mail"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("🔙 В НАЧАЛО"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Нажмите на кнопку")

// Kb_user_Hi Kb_user_Hi Kb_user_test 1234567
//var Kb_user_Hi = tu.Keyboard(
//	tu.KeyboardRow( // Row 1
//		tu.KeyboardButton("🛠 Работы / Услуги"),                                                // Column 1
//		tu.KeyboardButton("🔝 Главное меню").WithWebApp(tu.WebAppInfo("https://statutbot.ru")), // Column 2
//	),
//	tu.KeyboardRow( // Row 2
//		tu.KeyboardButton("Contact").WithRequestContact(),   // Column 1
//		tu.KeyboardButton("Location").WithRequestLocation(), // Column 2
//	),
//	tu.KeyboardRow( // Row 3
//		tu.KeyboardButton("Poll Any").WithRequestPoll(tu.PollTypeAny()),         // Column 1
//		tu.KeyboardButton("Poll Regular").WithRequestPoll(tu.PollTypeRegular()), // Column 2
//		tu.KeyboardButton("Poll Quiz").WithRequestPoll(tu.PollTypeQuiz()),       // Column 3
//	),
//).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
