package controller

import (
	//"fmt"
	"net/http"
	"net/url"
	"../utils"
	"../entity"

	"log"
)


//make order through transaction
func MakeOrder(w http.ResponseWriter, r *http.Request) {

	q,_ := url.ParseQuery(r.URL.Path)
	projectOwner := q["/submitOrder/pid"][0]
	investor := q["uname"][0]
	amount := q["amount"][0]

	utils.MakeTransaction(investor,projectOwner,amount + " USDT","eosio.token","transfer")
	log.Printf("after transfer USTD")
	utils.MakeTransaction(projectOwner,investor,amount + " ABC","eosio.token","transfer")
	log.Printf("after transfer ABC")

	var result entity.BalanceResult
	result.Ret = "200"

	var data entity.Data
	data.Code = "200"
	data.MSG ="OK"
	result.Data = data
	contentString := utils.StructToJson(result)

	utils.Output(w,contentString)

}
