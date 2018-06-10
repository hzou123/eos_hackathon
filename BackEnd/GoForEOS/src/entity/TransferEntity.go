package entity


type TransferEntity struct {
	Code   string `json:"code"`
	Action string `json:"action"`
	Args   struct {
		From     string `json:"from"`
		To       string `json:"to"`
		Quantity string    `json:"quantity"`
		Memo string `json:"memo"`
	} `json:"args"`
}