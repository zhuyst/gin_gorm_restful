package user_dao

import(
	".."
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email string
}

var db *gorm.DB

func init()  {
	db = model.GetDB()
	db.AutoMigrate(&User{})
}

func AddUser(user *User) error {
	return db.Create(user).Error
}

func UpdateUser(user *User) error {
	return db.Model(&User{}).Updates(user).Error
}

func GetUser(id uint) (user User,err error) {
	err = db.First(&user,id).Error
	return user,err
}

func ListUsers() (users []User,err error) {
	err = db.Find(&users).Error
	return users,err
}

func CountUsers() (count uint,err error) {
	err = db.Model(&User{}).Count(&count).Error
	return count,err
}

func DeleteUserByID(id uint) (err error) {
	user := User{}
	user.ID = id
	return DeleteUser(&user)
}

func DeleteUser(user *User) (err error) {
	return db.Delete(user).Error
}