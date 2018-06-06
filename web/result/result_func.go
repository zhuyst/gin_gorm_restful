package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetResultFunc func(c *gin.Context) (resultEntity Result)

func (getResult GetResultFunc) ToGinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		resultEntity := getResult(c)
		c.JSON(http.StatusOK,resultEntity.ToGinH())
	}
}
