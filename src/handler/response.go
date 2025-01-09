package handler

// RespBody response body.
type RespBody struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Token   string      `json:"token"`
}

// TrMsg translate the reason cause as a message

// NewRespBody new response body
func NewRespBody(code int, reason string) *RespBody {
	return &RespBody{
		Code:   code,
		Reason: reason,
	}
}

// NewRespBodyFromError new response body from error
func NewRespBodyFromError(e *CustomError) *RespBody {
	return &RespBody{
		Code:    e.Code,
		Reason:  e.Reason,
		Message: e.Message,
	}
}

// NewRespBodyData new response body with data
func NewRespBodyData(code int, reason string, data interface{}, token string) *RespBody {
	return &RespBody{
		Code:   code,
		Reason: reason,
		Result: data,
		Token:  token,
	}
}
