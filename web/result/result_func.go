package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取结果对象方法
type GetDataFunc func(c *gin.Context) (data interface{},err error)

// GetDataFunc -> Gin.HandlerFunc
func (getData GetDataFunc) ToGinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data,err := getData(c)
		resultEntity := BuildResult(data,err)
		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}
}
