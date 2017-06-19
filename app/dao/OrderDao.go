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


func FindOrderById (id int) model.OrderInfo{
	fmt.Println("id=", id)
	db := connect.Gdb
	fmt.Println("id=", id, "db=", db)
	tx := db.Begin()
	var orderInfo model.OrderInfo
	db.First(&orderInfo, id)
	tx.Commit()
	db.Close()
	fmt.Println("orderInfo=", orderInfo, ", db.Close()")
	return orderInfo
}

func FindOrderAll () []model.OrderInfo{
	db := connect.Gdb
	fmt.Println("db=", db)
	tx := db.Begin()
	var orderInfoList []model.OrderInfo
	db.Find(&orderInfoList)
	tx.Commit()
	db.Close()
	fmt.Println("orderInfoList=", orderInfoList, ", db.Close()")
	return orderInfoList
}

func UpdateOrderInfo(modelInfo model.OrderInfo) {
	fmt.Println("modelInfo=", modelInfo)
	db := connect.Gdb
	fmt.Println("modelInfo=", modelInfo, "db=", db)
	tx := db.Begin()

	if err := tx.Save(&modelInfo).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	db.Close()
	fmt.Println("modelInfo=", modelInfo, ", db.Close()")
}


func DeleteOrderById(id int) {
	fmt.Println("id=", id)
	db := connect.Gdb
	fmt.Println("id=", id, "db=", db)
	tx := db.Begin()
	var queryInfo model.OrderInfo
	db.First(&queryInfo, id)
	fmt.Println("id=", id, "queryInfo=", queryInfo)
	if err := tx.Delete(&queryInfo).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	db.Close()
	fmt.Println("id=", id, ", db.Close()")
}

























