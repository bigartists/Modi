package controllers

import (
	"github.com/gin-gonic/gin"
	"modi/src/vars"
	"sync"
)

type JSONResult struct {
	Message string      `json:"message"`
	Code    int8        `json:"code"`
	Result  interface{} `json:"result"`
	Token   string      `json:"token"`
}

func NewJSONResult(result interface{}) *JSONResult {
	return &JSONResult{
		Message: "",
		Code:    0,
		Result:  result,
		Token:   "",
	}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult(nil)
		},
	}
}

type Output1 func(c *gin.Context, v interface{}) interface{}

type ResultFunc1 func(result interface{}, message string) func(output Output1) interface{}

func ResultWrapper1(c *gin.Context) ResultFunc1 {
	return func(result interface{}, message string) func(output Output1) interface{} {
		r := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(r)
		r.Message = message
		//r.Code = code
		token := c.GetString("token")
		r.Token = token
		r.Result = result

		//r.Result = map[string]interface{}{
		//	"data":  result,
		//	"token": token,
		//}

		return func(output Output1) interface{} {
			return output(c, r)
		}
	}
}

func OK1(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = vars.HTTPSUCCESS
		r.Message = vars.HTTPMESSAGESUCCESS
		return r
	}
	return nil
}

func Created1(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = vars.HTTPSUCCESS
		r.Message = vars.HTTPMESSAGESUCCESS
		return r
	}
	return nil
}

func Error1(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = vars.HTTPFAIL
		if r.Message == "" {
			r.Message = vars.HTTPMESSAGEFAIL
		}
		return r
	}
	return nil
}

func Unauthorized1(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = vars.HTTPUNAUTHORIZED
		if r.Message == "" {
			r.Message = vars.HTTPMESSAGEUNAUTHORIZED
		}
		return r
	}
	return nil
}
