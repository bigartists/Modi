package controllers

import (
	"github.com/bigartists/Modi/src/vars"
	"github.com/gin-gonic/gin"
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

type Output func(c *gin.Context, v interface{}) interface{}

type ResultFunc func(result interface{}, message string) func(output Output) interface{}

func ResultWrapper(c *gin.Context) ResultFunc {
	return func(result interface{}, message string) func(output Output) interface{} {
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

		return func(output Output) interface{} {
			return output(c, r)
		}
	}
}

func OK(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = vars.HTTPSUCCESS
		r.Message = vars.HTTPMESSAGESUCCESS
		return r
	}
	return nil
}

func Created(c *gin.Context, v interface{}) interface{} {
	// 将v 转成 *JSONResult 类型
	if r, ok := v.(*JSONResult); ok {
		r.Code = vars.HTTPSUCCESS
		r.Message = vars.HTTPMESSAGESUCCESS
		return r
	}
	return nil
}

func Error(c *gin.Context, v interface{}) interface{} {
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

func Unauthorized(c *gin.Context, v interface{}) interface{} {
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
