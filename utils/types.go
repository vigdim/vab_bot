package utils

type TgUser struct {
	TgId      int64
	FirstName string
	LastName  string
	UserName  string
	Language  string
}

// PayData Платежные данные
type PayData struct {
	Type   string `json:"type"`
	Event  string `json:"event"`
	Object struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Amount struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"amount"`
	} `json:"object"`
}
