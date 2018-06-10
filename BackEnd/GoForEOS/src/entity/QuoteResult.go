package entity

type QuoteResult struct {
	Ret string `json:ret`
	Data QuoteData `json:data`
}

type QuoteData struct {
	Code string `json:code`
	MSG string    `json:msg`
	Info QuoteINFO `json:info`
}

type QuoteINFO struct {
	Quote string `json:quote`
	Minimal string `json:minimal`
	Balance string `json:balance`
	Percentage string `json:percentage`
}
