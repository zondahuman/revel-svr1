package tests

import (
	"testing"
	"fmt"
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
	httpUrl := "http://localhost:9000/FindOrderAll"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}

func Test_UpdateOrderInfo(t *testing.T) {
	httpUrl := "http://localhost:9000/UpdateOrderInfo"
	request := make(map[string]string)
	request["id"] = "72"
	request["name"] = "abin72"
	request["age"] = "72"
	request["version"] = "72"

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	result := common.HttpPostForm(httpUrl,request, header)
	fmt.Println("result=", result)
}


func Test_DeleteOrderById(t *testing.T) {
	httpUrl := "http://localhost:9000/DeleteOrderById"
	request := make(map[string]string)
	request["id"] = "74"

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	result := common.HttpPostForm(httpUrl,request, header)
	fmt.Println("result=", result)
}






















