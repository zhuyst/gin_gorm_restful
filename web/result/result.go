package result

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int
	Message string
	Data interface{}
}

func (result *Result) ToGinH() (h gin.H) {
	h = gin.H{}

	h["code"] = result.Code
	h["message"] = result.Message
	h["data"] = result.Data

	return h
}

func GetResult(Data interface{},err error) (result Result) {
	if err == nil{
		result = Result{
			Code : http.StatusOK,
			Message : "OK",
			Data : Data,
		}
	}else{
		result = Result{
			Code : http.StatusInternalServerError,
			Message : err.Error(),
		}
	}

	return result
}
