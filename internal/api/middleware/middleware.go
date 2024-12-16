package middleware

import (
	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
)

const SessionKey = "sess"

func AuthEchoMiddleware(service domain.SessionService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			sess, err := service.CheckSession(context, context.Request().Header)
			if err != nil {
				return context.NoContent(401)
			}

			context.Set(SessionKey, sess)
			return next(context)
		}
	}
}
