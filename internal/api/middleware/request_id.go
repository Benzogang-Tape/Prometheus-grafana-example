package middleware

import (
	"sync/atomic"

	"github.com/labstack/echo"
)

const RequestIDKey = "req_id"

func RequestIDMiddleware() echo.MiddlewareFunc {
	var reqID uint64
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			context.Set(RequestIDKey, atomic.AddUint64(&reqID, 1))
			return next(context)
		}
	}
}
