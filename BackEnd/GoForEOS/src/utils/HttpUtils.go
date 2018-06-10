package utils

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"bytes"
	"fmt"
)


func GetContent(url string) string{
	resp, err := http.Get(url)

	if err != nil || resp == nil{
		log.Printf("%s url connection issue",url)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}
	pageContent := string(body)
	status := resp.StatusCode
	//返回的状态码

	if status == 200 {
		log.Println(pageContent)
	}else {
		log.Printf("%s url connection issue, status code:%d",url,status)
	}

	return pageContent

}

func PostHttp(url,para string) string {


	//json序列化
	post := para

	//fmt.Println(url, "post", post)

	var jsonStr = []byte(post)
	//fmt.Println("jsonStr", jsonStr)
	fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)

		return ""
	}
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return string(body)
}

func PostContent(url,para string) string{
	resp, err := http.Post(url,"application/json",strings.NewReader(para))

	if err != nil || resp == nil{
		log.Printf("%s url connection issue",url)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}
	pageContent := string(body)
	status := resp.StatusCode
	//返回的状态码

	if status == 200 {
		log.Println(pageContent)
	}else {
		log.Printf("%s url connection issue, status code:%d",url,status)
	}

	return pageContent

}
