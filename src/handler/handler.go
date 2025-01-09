package handler

import (
	"errors"
	"github.com/bigartists/Modi/src/reason"
	"github.com/gin-gonic/gin"

	"log/slog"
	"net/http"
	"runtime"
)

// HandleResponse Handle response body
func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	token := ctx.GetString("token")
	if err == nil {
		ctx.JSON(http.StatusOK, NewRespBodyData(http.StatusOK, reason.Success, data, token))
		return
	}

	var customErr *CustomError
	// unknown error
	if !errors.As(err, &customErr) {
		buf := make([]byte, 1024)
		n := runtime.Stack(buf, false)
		slog.Error("Unknown error", "error", err, "stack", string(buf[:n]))
		ctx.JSON(http.StatusInternalServerError, NewRespBody(
			http.StatusInternalServerError, reason.UnknownError))
		return
	}

	// log internal server error
	if IsInternalServer(customErr) {
		slog.Error("Internal server error", "error", customErr)
	}
	respBody := NewRespBodyFromError(customErr)
	if data != nil {
		respBody.Result = data
	}
	ctx.JSON(customErr.Code, respBody)
}

// BindAndCheck bind request and check
func BindAndCheck(ctx *gin.Context, data interface{}) bool {

	if err := ctx.ShouldBind(data); err != nil {
		slog.Error("Bind request failed", "error", err)
		err := NewCustomError().WithCode(http.StatusBadRequest).WithReason(reason.RequestFormatError).WithMsg(err.Error())
		HandleResponse(ctx, err, nil)
		return true
	}
	return false
}
