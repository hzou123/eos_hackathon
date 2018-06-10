package entity

import (
	"strings"
	"log"
)

type TableRowEntity struct {
	More bool `json:"more"`
	Rows []struct {
		Balance string `json:"balance"`
	} `json:"rows"`
}



func (tr *TableRowEntity) GetBalance() [2]ASSET{
	asset := [2]ASSET{}
	l :=  len(tr.Rows)
	log.Printf("length:%d",l)
	var i = 0
	for i=0;i<l;i++ {
		s := tr.Rows[i].Balance
		o := strings.Split(s," ")
		asset[i].Quantity = o[0]
		asset[i].Token = o[1]
		log.Printf("quantity:%s",asset[i].Quantity)
		log.Printf("token:%s",asset[i].Token)



	}

	return asset
}
