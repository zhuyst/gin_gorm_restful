package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuyst/gin_gorm_restful/web/user-controller"
)

var router *gin.Engine

func init()  {
	router = gin.Default()
}

// 启动Gin
func Run()  {
	user_controller.SetRouterGroup(router)
	router.Run()
}
