package entity

type BINEntity struct {
	Binargs       string        `json:"binargs"`
	RequiredScope []interface{} `json:"required_scope"`
	RequiredAuth  []interface{} `json:"required_auth"`
}