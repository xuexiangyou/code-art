package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code 	int 			`json:"code"`
	Msg 	string			`json:"msg"`
	Data 	interface{}		`json:"data"`
	Meta 	Meta 			`json:"meta"`

	Errors  []ErrorItem 	`json:"errors"`
}

type Meta struct {
	RequestId 	string 		`json:"requestId"`
}

type ErrorItem struct {
	Key 		string		`json:"key"`
	Value 		string		`json:"value"`
}

func NewResponse() *Response {
	return &Response{
		Code: SUCCESS,
		Msg:    "",
		Data:       nil,
		Meta: Meta{
			RequestId: "",
		},
		Errors: []ErrorItem{},
	}
}

// Wrapper include context
type Wrapper struct {
	ctx *gin.Context
}

// WrapContext
func WrapContext(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx:ctx}
}

// Json 输出json,支持自定义response结构体
func (wrapper *Wrapper) Json(httpCode int, response *Response) {
	wrapper.ctx.JSON(httpCode, response)
}

// Success 成功的输出
func (wrapper *Wrapper) Success( data interface{}) {
	response := NewResponse()
	response.Data = data
	wrapper.Json(http.StatusOK,response)
}

// Error 错误输出
func (wrapper *Wrapper) Error(httpCode int, statusCode int, err ...interface{}) {
	response := NewResponse()
	response.Code = statusCode
	response.Msg = GetMsg(statusCode, err...)
	wrapper.Json(httpCode, response)
}
