package tests

import (
	"testing"
	"fmt"
	"encoding/json"
	"revel-svr1/app/util"
)

func Test_OrderAdd(t *testing.T) {
	httpUrl := "http://localhost:9000/OrderAdd"
	request := make(map[string]string)
	request["Id"] = "1"
	request["name"] = "abin1"
	request["age"] = "24"

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	result := common.HttpPostForm(httpUrl,request, header)
	//result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}

func Test_FindOrderById(t *testing.T) {
	httpUrl := "http://localhost:9000/FindOrderById?id=72"

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_FindOrderAll(t *testing.T) {
	httpUrl := "http://localhost:8100/order/FindOrderAll?id=59"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}

func Test_UpdateOrder(t *testing.T) {
	httpUrl := "http://localhost:8100/order/UpdateOrder"
	request := make(map[string]interface{})
	request["Id"] = 64
	request["Name"] = "abin29"
	request["Age"] = 29
	request["Version"] = 5

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}



func Test_Bool1(t *testing.T) {
	var flag bool
	flag = true
	if flag {
		fmt.Println("flag====true")
	}else{
		fmt.Println("flag====false")
	}
	fmt.Println("flag=", flag)
}

func Test_Bool2(t *testing.T) {
	var flag bool
	flag = true
	if flag==true {
		fmt.Println("flag====true")
	}else{
		fmt.Println("flag====false")
	}
	fmt.Println("flag=", flag)
}


















