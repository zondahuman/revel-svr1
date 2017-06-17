package service

import (
	"time"
	"revel-svr1/app/model"
	"revel-svr1/app/dao"
)

func AddOrderInfo(name string, age int)  {
	currentTime := time.Now()
	modelInfo  := model.OrderInfo{Name:name, Age:age, CreateTime:currentTime, UpdateTime:currentTime, Version:0}
	dao.AddOrderInfo(modelInfo)
}
