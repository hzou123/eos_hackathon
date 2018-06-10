package controller

import (
	"net/http"
	"net/url"
	"../entity"
	"../utils"
	"log"
)


//get balance of account
func GetUserBalance(w http.ResponseWriter, r *http.Request) {

	//get para
	q,_ := url.ParseQuery(r.URL.Path)
	uname := q["/getUserBalance/uname"][0]

	log.Println(uname)
	if uname == "" {
		log.Println("has no uname")
		return
	}

	//set result content
	var result entity.BalanceResult
	result.Ret = "200"

	var data entity.Data
	data.Code = "200"
	data.MSG ="OK"

	//get page content
	tableRowEntity := utils.GetTableRows(uname)
	if len(tableRowEntity.GetBalance()) == 0 {
		data.ASSETS = entity.Data{}.ASSETS
	}

	data.ASSETS = tableRowEntity.GetBalance()

	result.Data = data

	//to json format
	content := utils.StructToJson(result)
	utils.Output(w,content)
}
