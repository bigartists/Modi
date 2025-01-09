package handler

type CustomError struct {
	Code    int
	Reason  string
	Message string
}

func NewCustomError() *CustomError {
	return &CustomError{}
}

func (e *CustomError) Error() string {
	return e.Reason
}

func (e *CustomError) BadRequest(reason string) *CustomError {
	return e.WithCode(400).WithReason(reason)
}

func (e *CustomError) WithMsg(msg string) *CustomError {
	e.Message = msg
	return e
}

func (e *CustomError) WithCode(code int) *CustomError {
	e.Code = code
	return e
}

func (e *CustomError) WithReason(reason string) *CustomError {
	e.Reason = reason
	return e
}

func IsInternalServer(err *CustomError) bool {
	return err.Code >= 500
}
