package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuyst/gin_gorm_restful/model/user-dao"
	"github.com/zhuyst/gin_gorm_restful/util"
	"github.com/zhuyst/gin_gorm_restful/web/result"
)

var(
	getUser result.GetDataFunc    // 获取用户Handler
	listUsers result.GetDataFunc  // 获取用户列表Handler
	addUser result.GetDataFunc    // 新增用户Handler
	updateUser result.GetDataFunc // 更新用户Handler
	deleteUser result.GetDataFunc // 删除用户Handler
)

func init(){
	InitRouterHandler()
}

// 设置用户路由组
func SetRouterGroup(router *gin.Engine){
	userGroup := router.Group("/users")
	userGroup.GET("/:id",getUser.ToGinHandler()).
		GET("/",listUsers.ToGinHandler()).
		POST("/",addUser.ToGinHandler()).
		PUT("/:id",updateUser.ToGinHandler()).
		DELETE("/:id",deleteUser.ToGinHandler())
}

// 初始化路由Handler
func InitRouterHandler()  {
	getUser = func(c *gin.Context) (data interface{},err error) {
		// 从URL中获取用户ID
		id,err := util.Str2Uint(c.Param("id"))
		if err != nil{
			return nil,err
		}

		user,err := user_dao.GetUser(id)
		return user,err
	}

	listUsers = func(c *gin.Context) (data interface{},err error) {
		users,err := user_dao.ListUsers()
		return users,err
	}

	addUser = func(c *gin.Context) (data interface{},err error) {
		user,err := GetUserByContext(c)

		if err != nil {
			return nil, err
		}

		err = user_dao.AddUser(&user)
		return user,err
	}

	updateUser = func(c *gin.Context) (data interface{},err error) {
		user,err := GetUserByContext(c)

		if err != nil {
			return nil, err
		}

		// 从URL中获取用户ID
		id,err := util.Str2Uint(c.Param("id"))
		if err != nil {
			return nil,err
		}

		user.ID = id

		err = user_dao.UpdateUser(&user)
		return user,err
	}

	deleteUser = func(c *gin.Context) (data interface{},err error) {
		// 从URL中获取用户ID
		id,err := util.Str2Uint(c.Param("id"))
		if err != nil{
			return nil,err
		}

		err = user_dao.DeleteUserByID(id)
		return nil,err
	}
}

func GetUserByContext(c *gin.Context) (user user_dao.User,err error) {
	user = user_dao.User{}

	// 将JSON转为User对象
	err = c.ShouldBindJSON(&user)

	return user,err
}