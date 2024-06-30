package site_router

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/valyala/fasthttp"
	"github.com/vigdim/vab_bot/database/methods"
	"github.com/vigdim/vab_bot/keyboards"
	bot_router "github.com/vigdim/vab_bot/router/bot"
	"github.com/vigdim/vab_bot/utils"
	"io"
	"log"
	"strconv"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
	fmt.Println("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func AccountGet(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/account.html")
}

func AccountPost(ctx *fasthttp.RequestCtx) {
	sName := string(ctx.FormValue("name"))
	sPhone := string(ctx.FormValue("phone"))
	sEmail := string(ctx.FormValue("email"))
	sUserId := string(ctx.FormValue("us_id"))

	// Создаем или обновляем пользователя. При обновлении роль не изменяется
	err := methods.AddUser(sUserId, sEmail, sName, sPhone, 0, "user")
	if err != nil {
		ctx.Error(err.Error(), 0)
	}
	us_id_i64, _ := strconv.ParseInt(sUserId, 10, 64)
	var bot = bot_router.Bot

	_, _ = bot.SendMessage(tu.Message(tu.ID(us_id_i64),
		fmt.Sprintf("<b>Сохранённые данные:</b>\nИмя: <b>%s</b>\nEmail: <b>%s</b>\nТелефон: <b>%s</b>",
			sName, sEmail, sPhone)).WithReplyMarkup(keyboards.Kb_сabinet).WithParseMode(telego.ModeHTML))

	// MessFlagForDel - подтверждаем флагом удаление предыдущего сообщения и удаляем его
	// По сути удаляется сообщение с кнопкой 📰 Внести свои данные, так как данные внутри неё остались старые
	utils.MessFlagForDel = true
	utils.DelMessage(bot, utils.CurUpdate)
}

func Assets(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("views/pages/assets/js/inits.js")
}

func Pay(ctx *fasthttp.RequestCtx) {
	payBody := ctx.PostBody()
	fmt.Println(string(payBody)) // this is fine

	payBodyData := utils.PayData{} // Структура патежа

	err := json.Unmarshal(payBody, &payBodyData)
	if err != nil {
		return
	}

	fmt.Println(payBodyData.Object.ID)
	fmt.Println(payBodyData.Event)
	fmt.Println(payBodyData.Object.Status)
	fmt.Println(payBodyData.Object.Amount.Value)
	fmt.Println(payBodyData.Object.Amount.Currency)

	payload1 := `{"app_id":"20503546","command":{"c_num":"1719407813348","goods":[{"sum":120,"name":"Ваш любимый товар 1","count":1,"price":120,"nds_value":20,"nds_not_apply":false,"item_type":1,"payment_mode":4}],"tag1055":1,"author":"Кассир","smsEmail54FZ":"example@client.ru","payed_cash":0,"payed_prepay":0,"payed_cashless":120,"payed_nextcredit":0,"payed_consideration":0},"nonce":"salt_1719407813348","type":"printCheck"}`
	payload2 := "z2lzJz"
	h := md5.New()
	io.WriteString(h, payload1)
	io.WriteString(h, payload2)
	fmt.Printf("%x", h.Sum(nil))

}

func Init() {
	r := router.New()

	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/account", AccountGet)
	r.POST("/account_post", AccountPost)
	r.GET("/assets/js/inits.js", Assets)
	r.POST("/pay", Pay)

	// ngrok http --domain=workable-grouse-clean.ngrok-free.app 80
	// https://workable-grouse-clean.ngrok-free.app/
	log.Fatal(fasthttp.ListenAndServe(":80", r.Handler))
}
