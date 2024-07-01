package utils

type TgUser struct {
	TgId      int64
	FirstName string
	LastName  string
	UserName  string
	Language  string
}

// PayData Платежные данные для отравки
type PayData struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Amount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Description string `json:"description"`
	Recipient   struct {
		AccountId string `json:"account_id"`
		GatewayId string `json:"gateway_id"`
	}
	CreatedAt    string `json:"created_at"`
	Confirmation struct {
		Type            string `json:"type"`
		ConfirmationUrl string `json:"confirmation_url"`
	}
	Test       bool `json:"test"`
	Paid       bool `json:"paid"`
	Refundable bool `json:"refundable"`
	Metadata   struct{}
}

// PayDataCallback Прием ответа о платеже (от структуры PayData)
type PayDataCallback struct {
	Type   string `json:"type"`
	Event  string `json:"event"`
	Object struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Amount struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"amount"`
		IncomeAmount struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		}
		Description string `json:"description"`
		Recipient   struct {
			AccountId string `json:"account_id"`
			GatewayId string `json:"gateway_id"`
		}
		PaymentMethod struct {
			Type          string `json:"type"`
			Id            string `json:"id"`
			Saved         bool   `json:"saved"`
			Title         string `json:"title"`
			AccountNumber string `json:"account_number"`
		}
		CaptureAt      string `json:"capture_at"`
		CreatedAt      string `json:"created_at"`
		Test           bool   `json:"test"`
		RefundedAmount struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		}
		Paid       bool `json:"paid"`
		Refundable bool `json:"refundable"`
		Metadata   struct{}
	} `json:"object"`
}
