package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../result"
	"../../util"
	"../../model/user-dao"
)

var getUser func(c *gin.Context)
var listUsers func(c *gin.Context)
var addUser func(c *gin.Context)
var updateUser func(c *gin.Context)
var deleteUser func(c *gin.Context)

func init(){
	InitRouterHandler()
}

func SetRouterGroup(router *gin.Engine){
	userGroup := router.Group("/users")
	userGroup.GET("/:id",getUser).
		GET("/",listUsers).
		POST("/",addUser).
		PUT("/:id",updateUser).
		DELETE("/:id",deleteUser)
}

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
		user := user_dao.User{}
		err := c.ShouldBindJSON(&user)

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
		user := user_dao.User{}
		err := c.ShouldBindJSON(&user)

		var resultEntity result.Result
		if err != nil {
			resultEntity = result.GetResult(nil, err)
		}else{
			id,_ := util.Str2Uint(c.Param("id"))
			user.ID = id

			err = user_dao.UpdateUser(&user)
			resultEntity = result.GetResult(user,err)
		}

		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}

	deleteUser = func(c *gin.Context){
		id,_ := util.Str2Uint(c.Param("id"))
		err := user_dao.DeleteUserByID(id)

		resultEntity := result.GetResult(nil,err)
		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}
}