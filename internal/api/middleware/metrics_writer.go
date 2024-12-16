package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain/metrics"
)

func MetricsWriterMiddleware(m *metrics.NativeMetrics) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			start := time.Now()
			err := next(context)

			m.Timings.WithLabelValues(context.Request().URL.Path).Observe(time.Since(start).Seconds())
			m.Hits.WithLabelValues(strconv.Itoa(context.Response().Status), context.Request().URL.Path).Inc()

			return err
		}
	}
}
