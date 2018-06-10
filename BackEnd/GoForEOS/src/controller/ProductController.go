package controller

import (
	//"fmt"
	"net/http"
	"net/url"
	"../entity"
	"../utils"
	"log"
)

//show quote of product
func GetQuoteInfo(w http.ResponseWriter, r *http.Request)  {

	q,_ := url.ParseQuery(r.URL.Path)

	pname := q["/GetQuoteInfo/projectName"][0]
	uname := q["investor"][0]

	log.Println(pname + uname)

	//get account balance
	projectOwner := utils.GetTableRows(pname)
	investor := utils.GetTableRows(uname)

	var quote,balance string
	balanceEntity  := projectOwner.GetBalance()
	if len(balanceEntity) == 0 {
		log.Println("has no project info")
		quote = "0"
	}else
	{
		quote  = balanceEntity[0].Quantity
	}

	log.Println("quote:%s",quote)
	balanceEntity  = investor.GetBalance()

	if len(balanceEntity) == 0  {
		log.Println("has no investor info")
		balance = "0"
	}else {
		balance  = balanceEntity[0].Quantity
	}

	//set result content
	var result entity.QuoteResult
	result.Ret = "200"

	var data entity.QuoteData
	data.Code = "200"
	data.MSG ="OK"

	var quoteEntity entity.QuoteINFO
	quoteEntity.Balance = balance + " USTD"
	quoteEntity.Quote = quote + " USTD"
	quoteEntity.Minimal = "10000 USTD"
	quoteEntity.Percentage = "60%"
	data.Info = quoteEntity
	result.Data = data

	//to json format
	content := utils.StructToJson(result)
	utils.Output(w,content)

}
