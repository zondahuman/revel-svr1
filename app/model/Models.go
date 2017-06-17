package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type OrderInfo struct {
	Id         int  `primary key Id`
	Name       string  `Name`
	Age        int `Age`
	CreateTime time.Time `create_time`
	UpdateTime time.Time `update_time`
	Version    int  `version `
}

