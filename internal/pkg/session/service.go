package session

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/api/middleware"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain/metrics"
)

type service struct {
	mtr *metrics.ThirdPartyMetrics
}

func NewService(m *metrics.ThirdPartyMetrics) domain.SessionService {
	return service{
		mtr: m,
	}
}

func (s service) CheckSession(ctx echo.Context, headers http.Header) (domain.Session, error) {
	outerSource := "http://vk-golang.ru:17000/int/CheckSession?stable=true&light=true"
	source := "CheckSession"
	subject := "session"
	req, err := http.NewRequest(http.MethodGet, outerSource, nil)
	if err != nil {
		return domain.Session{}, err
	}

	req.Header = headers

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	timing := time.Since(start)

	if err != nil {
		ctx.Logger().Errorf("[%s] req_id:%d domain:%s src:%s endpoint:%s timing:%s err:%s",
			start.Format("2006-01-02 15:04:05"),
			ctx.Get(middleware.RequestIDKey).(uint64),
			subject,
			source,
			outerSource,
			timing.String(),
			err.Error(),
		)
		return domain.Session{}, err
	}

	ctx.Logger().Infof("[%s] req_id:%d domain:%s src:%s endpoint:%s timing:%s",
		start.Format("2006-01-02 15:04:05"),
		ctx.Get(middleware.RequestIDKey).(uint64),
		subject,
		source,
		outerSource,
		timing.String(),
	)

	s.mtr.Timings.WithLabelValues(outerSource).Observe(timing.Seconds())
	s.mtr.Hits.WithLabelValues(resp.Status, outerSource).Inc()

	switch resp.StatusCode {
	case 500:
		return domain.Session{}, errors.New("failed to request check session")
	case 200:
		return domain.Session{}, nil
	default:
		return domain.Session{}, domain.ErrNoSession
	}
}
