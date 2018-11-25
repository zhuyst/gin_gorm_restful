package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 数据库实例
var db *gorm.DB

const (
	SqlType = "mysql"       // 数据库类型
	Username = "test"       // 数据库用户名
	Password = "zhuyst"     // 数据库密码
	Hostname = "localhost"  // 数据库主机名
	Port = "3306"           // 数据库端口号
	Database = "test"       // 数据库名
)

func init() {
	var err error

	// 与数据库建立连接
	db, err = gorm.Open(SqlType, GetDbUrl())
	if err != nil {
		panic("数据库连接失败")
	}

	// 设置生成的数据表均为单数形式
	db.SingularTable(true)
}

// 获取数据库实例
func GetDB() *gorm.DB {
	return db
}

// 获取GORM的连接URL
func GetDbUrl() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		Username,Password,Hostname,Port,Database)
}