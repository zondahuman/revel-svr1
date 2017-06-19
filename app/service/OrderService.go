package service

import (
	"time"
	"revel-svr1/app/model"
	"revel-svr1/app/dao"
	"revel-svr1/app/pojo/base"
)

func AddOrderInfo(name string, age int)  {
	currentTime := time.Now()
	modelInfo  := model.OrderInfo{Name:name, Age:age, CreateTime:currentTime, UpdateTime:currentTime, Version:0}
	dao.AddOrderInfo(modelInfo)
}


func FindOrderById(id int) model.OrderInfo {
	modelInfo  := dao.FindOrderById(id)
	return modelInfo
}



func FindOrderAll() []model.OrderInfo {
	modelInfoList  := dao.FindOrderAll()
	return modelInfoList
}

func UpdateOrderInfo(id int, name string, age int, version int)  {
	currentTime := time.Now()
	modelInfo  := model.OrderInfo{Id:id, Name:name, Age:age, CreateTime:currentTime, UpdateTime:currentTime, Version:version}
	dao.UpdateOrderInfo(modelInfo)
}

func DeleteOrderById(id int) base.BaseResponse  {
	baseResponse := dao.DeleteOrderById(id)
	return baseResponse
}















