package user_controller

import (
	"github.com/gin-gonic/gin"
	"../result"
	"../../util"
	"../../model/user-dao"
)

var(
	getUser result.GetResultFunc     // 获取用户Handler
	listUsers result.GetResultFunc   // 获取用户列表Handler
	addUser result.GetResultFunc     // 新增用户Handler
	updateUser result.GetResultFunc  // 更新用户Handler
	deleteUser result.GetResultFunc  // 删除用户Handler
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
	getUser = func(c *gin.Context) (resultEntity result.Result) {
		// 从URL中获取用户ID
		id,err := util.Str2Uint(c.Param("id"))
		if err != nil{
			return result.GetResult(nil,err)
		}

		user,err := user_dao.GetUser(id)
		return result.GetResult(user,err)
	}

	listUsers = func(c *gin.Context) (resultEntity result.Result) {
		users,err := user_dao.ListUsers()
		return result.GetResult(users,err)
	}

	addUser = func(c *gin.Context) (resultEntity result.Result) {
		user,err := GetUserByContext(c)

		if err != nil {
			return result.GetResult(nil, err)
		}

		err = user_dao.AddUser(&user)
		return result.GetResult(user,err)
	}

	updateUser = func(c *gin.Context) (resultEntity result.Result) {
		user,err := GetUserByContext(c)

		if err != nil {
			return result.GetResult(nil, err)
		}

		// 从URL中获取用户ID
		id,err := util.Str2Uint(c.Param("id"))
		if err != nil {
			return result.GetResult(nil,err)
		}

		user.ID = id

		err = user_dao.UpdateUser(&user)
		return result.GetResult(user,err)
	}

	deleteUser = func(c *gin.Context) (resultEntity result.Result) {
		// 从URL中获取用户ID
		id,err := util.Str2Uint(c.Param("id"))
		if err != nil{
			return result.GetResult(nil,err)
		}

		err = user_dao.DeleteUserByID(id)
		return result.GetResult(nil,err)
	}
}

func GetUserByContext(c *gin.Context) (user user_dao.User,err error) {
	user = user_dao.User{}

	// 将JSON转为User对象
	err = c.ShouldBindJSON(&user)

	return user,err
}