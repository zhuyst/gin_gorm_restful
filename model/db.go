package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "test:zhuyst@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		panic("数据库连接失败")
	}

	db.SingularTable(true)
}

func GetDB() *gorm.DB {
	return db
}