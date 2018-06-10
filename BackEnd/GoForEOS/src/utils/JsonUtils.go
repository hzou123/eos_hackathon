package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"net/http"
)




// json to struct
func JsonToStruct(jsonStream string) interface{} {

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	var m interface{}

	for {
		err := dec.Decode(&m)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
		}
	}

	return m
}

// struct to json
func StructToJson(entity interface{}) string {
	bytes, _ := json.Marshal(entity)
	fmt.Printf("json:m,%s\n", bytes)

	return string(bytes)
}

//print out content
func Output(w http.ResponseWriter,content string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(content))
}