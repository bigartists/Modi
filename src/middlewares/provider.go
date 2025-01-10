package middlewares

import "github.com/google/wire"

var MiddlewareSets = wire.NewSet(
	NewAuthMiddleware,
	NewErrorHandlerMiddleware,
)
