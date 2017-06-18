package controllers

import (
	"github.com/revel/revel"
	"revel-svr1/app/service"
	"fmt"
	"encoding/json"
)



type OrderInfoController struct {
	*revel.Controller
	//Gtx *gorm.DB
}

/**
       http://localhost:9000/OrderAdd?name=lee&age=21
 */
func (this OrderInfoController) OrderAdd(name string, age int) revel.Result {
	fmt.Println("name=",name, "age=",age)
	service.AddOrderInfo(name, age)
	return this.RenderText("hi,"+name)
}

func (this OrderInfoController) FindOrderById(id int) revel.Result {
	fmt.Println("id=",id)
	modelInfo := service.FindOrderById(id)
	json, err := json.Marshal(modelInfo)
	fmt.Println("id=",id, "json=",json)
	if err != nil {
	    panic(err)
	}
	result := string(json)
	fmt.Println("id=",id, "result=",result)
	return this.RenderText(result)
}



func (this OrderInfoController) FindOrderAll(id int) revel.Result {
	fmt.Println("id=",id)
	modelInfoList := service.FindOrderAll()
	json, err := json.Marshal(modelInfoList)
	fmt.Println("id=",id, "json=",json)
	if err != nil {
	    panic(err)
	}
	result := string(json)
	fmt.Println("id=",id, "result=",result)
	return this.RenderText(result)
}




func (this OrderInfoController) UpdateOrderInfo(id int,name string, age int, version int) revel.Result {
	fmt.Println("id=",id, "name=",name, "age=",age)
	service.UpdateOrderInfo(id, name, age, version)
	return this.RenderText("hi,"+name)
}
































