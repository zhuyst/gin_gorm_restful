package web

import (
	"github.com/gin-gonic/gin"
	"./user_controller"
)

var router *gin.Engine

func init()  {
	router = gin.Default()
}

func Run()  {
	user_controller.SetRouterGroup(router)
	router.Run()
}
