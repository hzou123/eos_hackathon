package entity

type PushTransactionEntity struct {
	Signatures     []string `json:"signatures"`
	Compression 	string  `json:"compression"`
	ContextFreeData     []string `json:"context_free_data"`
	Transaction TransactionEntity `json:"transaction"`
}

type TransactionEntity struct {
	Region int  `json:"region"`
	RefBlockNum    string   `json:"ref_block_num"`
	RefBlockPrefix string   `json:"ref_block_prefix"`
	Expiration     string   `json:"expiration"`
	MacNetUsageWords     int   `json:"max_net_usage_words"`
	MaxKcpuUsage     int   `json:"max_kcpu_usage"`
	DelaySec     int   `json:"delay_sec"`
	ContextFreeActions []string `json:"context_free_actions"`
	Scope          [2]string `json:"scope"`
	Actions 	   [1]Action	`json:"actions"`

	//	Authorizations []string `json:"authorizations"`
	//PublicKey 	string `json:"public_key"`

}
type Action struct {
	Account          string   `json:"account"`
	Name          string   `json:"name"`
	Authorization [1]struct {
		Actor    string `json:"actor"`
		Permission string `json:"permission"`
	} `json:"authorization"`
	Data string `json:"data"`
}
