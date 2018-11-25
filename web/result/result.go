package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通用API返回模板类
type Result struct {
	Code int          // 状态码
	Message string    // 结果信息
	Data interface{}  // 返回内容实体
}

// 将Result对象转为Gin的结果对象
func (result *Result) ToGinH() (h gin.H) {
	h = gin.H{}

	h["code"] = result.Code
	h["message"] = result.Message
	h["data"] = result.Data

	return h
}

// 获取Result对象
func BuildResult(data interface{},err error) Result {
	if err != nil{
		return Result{
			Code : http.StatusInternalServerError,
			Message : err.Error(),
		}
	}

	return Result{
		Code : http.StatusOK,
		Message : "OK",
		Data : data,
	}
}