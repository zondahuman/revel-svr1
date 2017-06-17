package dao

import (
	"fmt"
	"revel-svr1/app/model"
	"revel-svr1/app/connect"
)

func AddOrderInfo1(modelInfo model.OrderInfo) {
	fmt.Println("modelInfo=", modelInfo)
	db := connect.Gdb
	tx := db.Begin()
	//defer func() {
	//	// 必须要先声明defer，否则不能捕获到panic异常
	//	fmt.Println("c---------------------------------------")
	//	if err := recover(); err != nil {
	//		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	//		fmt.Println("brfore --tx.Rollback---------------------------------------")
	//		tx.Rollback()
	//		fmt.Println("after --tx.Rollback---------------------------------------")
	//
	//	}
	//	fmt.Println("d--------------------------")
	//}()
	flag := db.NewRecord(modelInfo)
	fmt.Println("modelInfo=", modelInfo, ", flag", flag)
	if flag == true {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	db.Close()
	fmt.Println("modelInfo=", modelInfo, ", flag", flag, "db.Close()")

}

func AddOrderInfo(modelInfo model.OrderInfo) {
	fmt.Println("modelInfo=", modelInfo)
	db := connect.Gdb
	fmt.Println("modelInfo=", modelInfo, "db=", db)
	tx := db.Begin()

	flag := tx.NewRecord(modelInfo)
	fmt.Println("modelInfo=", modelInfo, ", flag", flag)

	if err := tx.Create(&modelInfo).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	db.Close()
	fmt.Println("modelInfo=", modelInfo, ", db.Close()")

}
