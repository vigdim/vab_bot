package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/vigdim/vab_bot/database/methods"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

func SendPayOfd(bot *telego.Bot, query telego.CallbackQuery, nameOfd string, OfdId string, namePeriod string,
	PeriodId string, namePrice string, PriceId string) {

	url := ACQUIRING_SERVER
	returnUrl := DOMAIN_SERVER + "/account"
	method := "POST"

	payload := strings.NewReader(`{
        "amount": {
          "value": "` + namePrice + `",
          "currency": "RUB"
        },
        "capture": true,
        "confirmation": {
          "type": "redirect",
          "return_url": "` + returnUrl + `"
        },
        "description": "Оплата кода активации для ОФД ` + nameOfd + `. Срок действия: ` + namePeriod + `."
      }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	IdempotenceKey, err := rand.Int(rand.Reader, big.NewInt(1000000000000))
	fmt.Println(IdempotenceKey)

	req.Header.Add("Idempotence-Key", fmt.Sprintf("%s", IdempotenceKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic NDA3NTUwOnRlc3RfSWlYX1RpX0x0UjlyZzZqZGtPS09jWnJMbXVyX3JWUlUxOGg1Z2RPaDJwcw==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	payBodyData := PayData{} // Структура патежа
	err = json.Unmarshal(body, &payBodyData)
	if err != nil {
		return
	}
	fmt.Println(payBodyData.Id)
	fmt.Println(payBodyData.Status)
	fmt.Println(payBodyData.Description)
	fmt.Println(payBodyData.Confirmation.ConfirmationUrl)
	fmt.Println(payBodyData.CreatedAt)

	TgID := Int64ToStr(query.Message.GetChat().ID)
	err = methods.GetSetPayStatus(TgID, payBodyData.Id, payBodyData.Status, nameOfd, OfdId, namePeriod, PeriodId, namePrice, PriceId)
	if err != nil {
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Ошибка при обращении к базе данных! Повторите позже!"))
		return
	}
	//exec.Command("open", payBodyData.Confirmation.ConfirmationUrl).Run()
}

func SendCheckOfd(nameOfd string, namePeriod string, namePrice string) {
	url := CHECK_SERVER
	method := "POST"

	payload1 := strings.NewReader(`{
    "app_id": "20503546",
    "command": {
        "c_num": "1719405133466",
        "goods": [
            {
                "sum": 150,
                "name": "Наименование товара",
                "count": 1,
                "price": 150,
                "nds_value": 20,
                "nds_not_apply": false,
                "item_type": 1,
                "payment_mode": 4
            }
        ],
        "tag1055": 1,
        "author": "Кассир",
        "smsEmail54FZ": "example@client.ru",
        "payed_cash": 0,
        "payed_prepay": 0,
        "payed_cashless": 150,
        "payed_nextcredit": 0,
        "payed_consideration": 0
    },
    "nonce": "salt_17333487",
    "type": "printCheck"
	}z2lzJz`)
	var b bytes.Buffer
	b.ReadFrom(payload1)
	data := []byte(b.String())
	fmt.Printf("%x", md5.Sum(data))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload1)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("sign", fmt.Sprintf("%x", md5.Sum(data)))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
