package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../result"
	"../../util"
	"../../model/user-dao"
)

// 获取用户Handler
var getUser func(c *gin.Context)

// 获取用户列表Handler
var listUsers func(c *gin.Context)

// 新增用户Handler
var addUser func(c *gin.Context)

// 更新用户Handler
var updateUser func(c *gin.Context)

// 删除用户Handler
var deleteUser func(c *gin.Context)

func init(){
	InitRouterHandler()
}

// 设置用户路由组
func SetRouterGroup(router *gin.Engine){
	userGroup := router.Group("/users")
	userGroup.GET("/:id",getUser).
		GET("/",listUsers).
		POST("/",addUser).
		PUT("/:id",updateUser).
		DELETE("/:id",deleteUser)
}

// 初始化路由Handler
func InitRouterHandler()  {
	getUser = func(c *gin.Context) {
		id,_ := util.Str2Uint(c.Param("id"))
		user,err := user_dao.GetUser(id)

		resultEntity := result.GetResult(user,err)
		c.JSON(http.StatusOK, resultEntity.ToGinH())
	}

	listUsers = func(c *gin.Context){
		users,err := user_dao.ListUsers()
		resultEntity := result.GetResult(users,err)
		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}

	addUser = func(c *gin.Context)  {
		user,err := GetUserByContext(c)

		var resultEntity result.Result
		if err != nil {
			resultEntity = result.GetResult(nil, err)
		}else{
			err = user_dao.AddUser(&user)
			resultEntity = result.GetResult(user,err)
		}

		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}

	updateUser = func(c *gin.Context) {
		user,err := GetUserByContext(c)

		var resultEntity result.Result
		if err != nil {
			resultEntity = result.GetResult(nil, err)
		}else{
			// 从URL中获取用户ID
			id,_ := util.Str2Uint(c.Param("id"))
			user.ID = id

			err = user_dao.UpdateUser(&user)
			resultEntity = result.GetResult(user,err)
		}

		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}

	deleteUser = func(c *gin.Context){

		// 从URL中获取用户ID
		id,_ := util.Str2Uint(c.Param("id"))
		err := user_dao.DeleteUserByID(id)

		resultEntity := result.GetResult(nil,err)
		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}
}

func GetUserByContext(c *gin.Context) (user user_dao.User,err error) {
	user = user_dao.User{}

	// 将JSON转为User对象
	err = c.ShouldBindJSON(&user)

	return user,err
}