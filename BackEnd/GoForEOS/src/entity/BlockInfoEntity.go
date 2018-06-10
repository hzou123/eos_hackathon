package entity

type BlockInfoEntity struct {
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Producer              string        `json:"producer"`
	ProducerChanges       []interface{} `json:"producer_changes"`
	ProducerSignature     string        `json:"producer_signature"`
	Cycles                []interface{} `json:"cycles"`
	ID                    string        `json:"id"`
	BlockNum              int           `json:"block_num"`
	RefBlockPrefix        int           `json:"ref_block_prefix"`
}