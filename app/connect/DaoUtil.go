package connect

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"revel-svr1/app/model"
)

var Gdb *gorm.DB
var err error


func init() {
	Gdb, err = gorm.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8&parseTime=true&loc=Local")
	//db, err := gorm.Open("mysql", "db:dbadmin@tcp(127.0.0.1:3306)/foo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	//defer Db.Close()
	Gdb.DB().SetMaxIdleConns(10)
	Gdb.DB().SetMaxOpenConns(100)
	Gdb.LogMode(true)
	// 自动迁移表，生成的表名为 products
	Gdb.Model(&model.OrderInfo{})
	//Gdb.AutoMigrate(&model.OrderInfo{})
	// Add table suffix when create tables
	//Gdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.OrderInfo{})
}
