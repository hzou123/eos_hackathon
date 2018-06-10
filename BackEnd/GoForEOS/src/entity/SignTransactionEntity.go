package entity

type SignTransactionEntity struct {
	RefBlockNum    string           `json:"ref_block_num"`
	RefBlockPrefix string         `json:"ref_block_prefix"`
	Expiration     string        `json:"expiration"`
	Scope          [2]string      `json:"scope"`
	ReadScope      []interface{} `json:"read_scope"`
	Messages       [1]Message `json:"messages"`
	Signatures []string `json:"signatures"`
}

type Message       struct {
	Code          string `json:"code"`
	Type          string `json:"type"`
	Authorization [1]struct {
		Actor    string `json:"actor"`
		Permission string `json:"permission"`
	} `json:"authorization"`
	Data string `json:"data"`
}


type SignResult struct {
	Expiration            string        `json:"expiration"`
	RefBlockNum           int           `json:"ref_block_num"`
	RefBlockPrefix        int64         `json:"ref_block_prefix"`
	MaxNetUsageWords      int           `json:"max_net_usage_words"`
	MaxCPUUsageMs         int           `json:"max_cpu_usage_ms"`
	DelaySec              int           `json:"delay_sec"`
	ContextFreeActions    []interface{} `json:"context_free_actions"`
	Actions               []interface{} `json:"actions"`
	TransactionExtensions []interface{} `json:"transaction_extensions"`
	Signatures            []string      `json:"signatures"`
	ContextFreeData       []interface{} `json:"context_free_data"`
}