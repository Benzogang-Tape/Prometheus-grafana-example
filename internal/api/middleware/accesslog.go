package middleware

import (
	"time"

	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
)

func AccessLogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			start := time.Now()
			err := next(context)
			timing := time.Since(start)

			if err != nil {
				context.Logger().Errorf("\n[%s] req_id:%d user_id:%s method:%s path:%s addr:%s status:%d duration:%s err:%s\n",
					start.Format("2006-01-02 15:04:05"),
					context.Get(RequestIDKey).(uint64),
					context.Get(SessionKey).(domain.Session).UserID,
					context.Request().Method,
					context.Request().URL.Path,
					context.Request().RemoteAddr,
					context.Response().Status,
					timing.String(),
					err.Error(),
				)
			} else {
				context.Logger().Infof("\n[%s] req_id:%d user_id:%s method:%s path:%s addr:%s status:%d duration:%s\n",
					start.Format("2006-01-02 15:04:05"),
					context.Get(RequestIDKey).(uint64),
					context.Get(SessionKey).(domain.Session).UserID,
					context.Request().Method,
					context.Request().URL.Path,
					context.Request().RemoteAddr,
					context.Response().Status,
					timing.String(),
				)
			}

			return err
		}
	}
}
