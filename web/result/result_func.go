package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取结果对象方法
type GetResultFunc func(c *gin.Context) (resultEntity Result)

// GetResultFunc -> Gin.HandlerFunc
func (getResult GetResultFunc) ToGinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		resultEntity := getResult(c)
		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}
}
