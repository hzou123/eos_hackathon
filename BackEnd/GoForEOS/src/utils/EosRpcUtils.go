package utils

import (
	"encoding/json"
	"fmt"
	"../entity"
	"strconv"
	"log"
	"time"
)

const GetChainUrl ="http://127.0.0.1:8888/v1/chain/get_info"
const GetBlockUrl="http://127.0.0.1:8888/v1/chain/get_block"
const JsonToBinUrl = "http://127.0.0.1:8888/v1/chain/abi_json_to_bin"
const PushTransactionUrl = "http://127.0.0.1:8888/v1/chain/push_transaction"
const UserAccountUrl = "http://127.0.0.1:8888/v1/chain/get_account"
const GetTableRowsUrl = "http://127.0.0.1:8888/v1/chain/get_table_rows"
const SignTransactionUrl = "http://localhost:8888/v1/wallet/sign_transaction"
const UnlockWalletUrl = "http://localhost:8888/v1/wallet/unlock"

//get chain info
func GetChain() *entity.ChainEntity{

	//get page content
	content := PostHttp(GetChainUrl,"")
	if content == "" {
		return nil
	}

	blockInfo := &entity.ChainEntity{}
	er := json.Unmarshal([]byte(content), &blockInfo)

	if(er != nil) {
		fmt.Println("Unmarshal faild")
	}else {
		fmt.Println("Account name:"+blockInfo.HeadBlockID)
	}

	return blockInfo
}

//get block info
func GetBlock(block_num_or_id int) *entity.BlockInfoEntity{

	//get page content
	content := PostHttp(GetBlockUrl,"{\"block_num_or_id\":\""+strconv.Itoa(block_num_or_id)+"\"}")
	if content == "" {
		return nil
	}

	blockInfo := &entity.BlockInfoEntity{}
	er := json.Unmarshal([]byte(content), &blockInfo)

	if(er != nil) {
		fmt.Println("Unmarshal faild")
	}else {
		fmt.Println("Account name:"+blockInfo.ID)
	}

	return blockInfo
}

//json to bin for transactoin
func JsonToBin(input string) string{
	//get page content
	content := PostHttp(JsonToBinUrl,input)
	if content == "" {
		return ""
	}

	log.Println("content:%s",content)
	binEntity := &entity.BINEntity{}
	er := json.Unmarshal([]byte(content), &binEntity)


	if(er != nil) {
		fmt.Println("Unmarshal faild")
	}else {
		fmt.Println("Account name:"+binEntity.Binargs)
	}

	return binEntity.Binargs
}

//get account info
func GetAccount(account string) *entity.AccountEntity{
	//get page content
	content := PostHttp(UserAccountUrl,"{\"account_name\":\""+account+"\"}")
	if content == "" {
		return nil
	}

	accountEntity := &entity.AccountEntity{}
	er := json.Unmarshal([]byte(content), &accountEntity)

	if(er != nil) {
		fmt.Println("Unmarshal faild")
	}else {
		fmt.Println("Account name:"+accountEntity.AccountName)
	}

	return accountEntity
}

//push transaction
func PushTransaction(transaction entity.PushTransactionEntity) {
	//get page content
	PostHttp(PushTransactionUrl,StructToJson(transaction))
}

//Get table rows
func GetTableRows(account string) entity.TableRowEntity {

	tableRows := entity.TableRowEntity{}

	//get page content
	content := PostHttp(GetTableRowsUrl,"{\"scope\":\""+account+"\", \"code\":\"eosio.token\", \"table\":\"accounts\", \"json\": true}")
	if content == "" {
		return tableRows
	}

	er := json.Unmarshal([]byte(content), &tableRows)

	if(er != nil) {
		fmt.Println("Unmarshal faild")
	}else {
	}

	return tableRows
}

//sign signature for transaction
func Sign(transaction entity.SignTransactionEntity,key string) entity.SignResult {
	sign := entity.SignResult{}
	var para string
	para = "["+StructToJson(transaction)+",[\""+key+"\"],\"\"]"
	content := PostHttp(SignTransactionUrl,para)


	json.Unmarshal([]byte(content), &sign)


	return sign
}

//unlock wallet
func UnlockWallet(uname,pwd string) {
	PostHttp(GetChainUrl,"[\""+uname+"\",\""+pwd+"\"]")
}

//make transaction
func MakeTransaction(from,to,amount,code,action string) {
	//get account balance
	investor := GetAccount(from)
	if investor == nil {
		log.Printf("has no investor info")
		return
	}

	//StructToJson(investor)
	key := investor.Permissions[0].RequiredAuth.Keys[0]

	transcation := entity.TransactionEntity{}
	signTransaction := entity.SignTransactionEntity{}

	chainEntity := GetChain()
	if chainEntity == nil {
		log.Printf("has no chain info")
		return
	}

	//get block info
	blockEntity := GetBlock(chainEntity.LastIrreversibleBlockNum)
	if blockEntity == nil {
		log.Printf("has no block info")
		return
	}
	transcation.RefBlockNum = strconv.Itoa(blockEntity.BlockNum)
	transcation.RefBlockPrefix = strconv.Itoa(blockEntity.RefBlockPrefix)
	signTransaction.RefBlockNum = transcation.RefBlockNum
	signTransaction.RefBlockPrefix = transcation.RefBlockPrefix


	//get json bin
	pushTransactoinEntity := entity.PushTransactionEntity{}
	transferEntity := entity.TransferEntity{}
	transferEntity.Code = code
	transferEntity.Action = action
	targs := entity.TransferEntity{}.Args
	targs.From = from
	targs.To = to
	targs.Quantity = amount
	transferEntity.Args = targs
	json := StructToJson(transferEntity)
	bin := JsonToBin(json)

	log.Println("after bin")

	//set action
	actions := [1]entity.Action{}
	actions[0].Account = transferEntity.Code
	actions[0].Data = bin
	actions[0].Name = transferEntity.Action

	/*	//set Recipients
		recps := actions[0].Recipients
		recps[0] = from
		recps[1] = to
		actions[0].Recipients = recps
	*/
	//set data
	actions[0].Data = bin

	messages := [1]entity.Message{}
	messages[0].Code = transferEntity.Code
	messages[0].Data = bin
	messages[0].Type = transferEntity.Action

	log.Println("after messages")

	//set authorization
	authorization := actions[0].Authorization
	authorization[0].Actor = from
	authorization[0].Permission = "active"
	actions[0].Authorization = authorization

	a := messages[0].Authorization
	a[0].Actor = from
	a[0].Permission = "active"
	messages[0].Authorization = a
	signTransaction.Messages = messages

	timestamp := time.Now().Unix()
	tm2 := time.Unix(timestamp, 0)
	expire := tm2.Format("2018-06-09T07:16:55")

	signTransaction.Expiration= expire
	signTransaction.Signatures = []string{}
	log.Println("after authorization")

	transcation.Actions = actions
	transcation.Region= 0
	transcation.MaxKcpuUsage = 0
	transcation.MacNetUsageWords = 0
	transcation.DelaySec = 0
	////as := transcation.Authorizations
	//transcation.Authorizations = as

	sig := []string{}
	pushTransactoinEntity.Signatures = sig
	pushTransactoinEntity.Compression = "none"
	pushTransactoinEntity.ContextFreeData = []string{}

	//transcation.Signatures = sig

	//auths := []string{}`
	//transcation.Authorizations = auths

	//set scope
	scope := transcation.Scope
	scope[0] = from
	scope[1] = to
	transcation.Scope = scope
	transcation.ContextFreeActions = []string{}
	signTransaction.Scope = scope

	//set public key
	//	transcation.PublicKey = key.Key

	//set Expiration
	transcation.Expiration = expire

	//sign transaction
	signEntity := Sign(signTransaction,key.Key)

	log.Printf("after sign")
	pushTransactoinEntity.Signatures = signEntity.Signatures
	//transcation.Actions[0].Authorization[0].Signatures = signEntity.Signatures
	pushTransactoinEntity.Transaction = transcation


	//push transaction
	PushTransaction(pushTransactoinEntity)

}


