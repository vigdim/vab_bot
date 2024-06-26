package utils

var (
	DOMAIN_SERVER    string
	ACQUIRING_SERVER string
	CHECK_SERVER     string
)

const (

	// ========================================================================
	WelcomeUserMessage  string = "<b>Здравствуйте, %s!\nДобро пожаловать в сервис ККТ от компании ВАБ!</b>"
	WelcomeAdminMessage string = "<b>Здравствуйте, %s!\nЗайдете под админом или пользователем?</b>"
	AboutCompany        string = "<b>© О компании</b>\n\n" +
		"<b>ООО «ВАБ» - «Внедрение Автоматизации Бизнеса»</b> зарегистрировано в 2017 году.\n\n" +
		"<b>Адрес:</b> 298403, Республика Крым, Бахчисарайский район, город Бахчисарай, ул. Ракитского, д. 11 офис 3\n\n" +
		"<b>ИНН / КПП:</b> 9104009766 / 910401001\n<b>ОГРН:</b> 1179102029963"

	MyDataMessage string = "<b>Твои данные:</b>\nTelegram id: <b><code>%s</code></b>\nFirst name: <b>%s</b>" +
		"\nLast name: <b>%s</b>\nUser name: <b>%s</b>\nLanguage code: <b>%s</b>"

	Consultation string = "<b>Бесплатно</b> проконсультируем Вас по всем вопросам, связанным с применением контрольно кассовой техники.\n" +
		"Можем предложить приобретение ККТ и полное либо частичное её сопровождение, а также бухгалтерское обслуживание организаций.\n\n" +
		"Звоните по телефону: <b>%s</b>\nИли воспользуйтесь кнопками ниже для обратной связи:"
	ConsultationPayKKT string = "❗️<i>Специалист нашей компании получит Ваше сообщение о консультации по приобретению ККТ.</i>\n" +
		"<b>Подтвердите отправку сообщения!</b> 👇"
	ConsultationPayKKTYes string = "❗️Консультацией по приобретению ККТ заинтересовался возможный клиент " +
		"<b>%s %s</b>. \n<b>Дайте ему ответ, щёлкнув по кнопке:</b>"
	ConsultationPayFnOfd string = "❗️<i>Специалист нашей компании получит Ваше сообщение о консультации по приобретению ФН / ОФД.</i>\n" +
		"<b>Подтвердите отправку сообщения!</b> 👇"
	ConsultationPayFnOfdYes string = "❗️Консультацией по приобретению ФН / ОФД заинтересовался возможный клиент " +
		"<b>%s %s</b>. \n<b>Дайте ему ответ, щёлкнув по кнопке:</b>"
	ConsultationServKKT string = "❗️<i>Специалист нашей компании получит Ваше сообщение о консультации по обслуживанию ККТ.</i>\n" +
		"<b>Подтвердите отправку сообщения!</b> 👇"
	ConsultationServKKTYes string = "❗️Консультацией по обслуживанию ККТ заинтересовался возможный клиент " +
		"<b>%s %s</b>. \n<b>Дайте ему ответ, щёлкнув по кнопке:</b>"
	ConsultationServBuh string = "❗️<i>Специалист нашей компании получит Ваше сообщение о консультации по бухгалтерскому обслуживанию.</i>\n" +
		"<b>Подтвердите отправку сообщения!</b> 👇"
	ConsultationServBuhYes string = "❗️Консультацией по бухгалтерскому обслуживанию заинтересовался возможный клиент " +
		"<b>%s %s</b>. \n<b>Дайте ему ответ, щёлкнув по кнопке:</b>"
	ConsultationMesSend string = "💬 <b>Сообщение отправлено!</b>\n<i>Наш специалист оперативно свяжется с Вами в рабочее время.</i> 😊"
)
