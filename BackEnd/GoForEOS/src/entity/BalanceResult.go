package entity

type BalanceResult struct {
	Ret string `json:ret`
	Data Data `json:data`
}

type Data struct {
	Code string `json:code`
	MSG string    `json:msg`
	ASSETS [2]ASSET `json:assets`
}
type ASSET struct {
	Token string    `json:token`
	Quantity    string    `json:quantity`
}