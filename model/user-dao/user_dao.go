package user_dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zhuyst/gin_gorm_restful/model"
)

type User struct {
	gorm.Model
	Username string  // 用户名
	Email string     // 邮箱
}

var db *gorm.DB

func init()  {
	db = model.GetDB()

	// 生成User表
	db.AutoMigrate(&User{})
}

// 新增用户
func AddUser(user *User) error {
	return db.Create(user).Error
}

// 更新用户
func UpdateUser(user *User) error {
	return db.Model(&User{}).Updates(user).Error
}

// 通过ID查询用户
func GetUser(id uint) (user User,err error) {
	err = db.First(&user,id).Error
	return user,err
}

// 查询用户列表
func ListUsers() (users []User,err error) {
	err = db.Find(&users).Error
	return users,err
}

// 查询用户记录数
func CountUsers() (count uint,err error) {
	err = db.Model(&User{}).Count(&count).Error
	return count,err
}

// 通过用户ID删除用户
func DeleteUserByID(id uint) (err error) {
	user := User{}
	user.ID = id
	return DeleteUser(&user)
}

// 删除用户
func DeleteUser(user *User) (err error) {
	return db.Delete(user).Error
}