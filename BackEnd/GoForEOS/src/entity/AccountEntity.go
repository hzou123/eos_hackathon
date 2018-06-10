package entity

type AccountEntity struct {
	AccountName string `json:"account_name"`
	Permissions []struct {
		PermName     string `json:"perm_name"`
		Parent       string `json:"parent"`
		RequiredAuth struct {
			Threshold int `json:"threshold"`
			Keys      []struct {
				Key    string `json:"key"`
				Weight int    `json:"weight"`
			} `json:"keys"`
			Accounts []interface{} `json:"accounts"`
		} `json:"required_auth"`
	} `json:"permissions"`
}