package controllers

import (
	"github.com/revel/revel"
	"revel-svr1/app/service"
	"fmt"
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










