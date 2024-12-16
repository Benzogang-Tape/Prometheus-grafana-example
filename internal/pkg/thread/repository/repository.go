package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/api/middleware"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain/metrics"
)

type repository struct {
	mtr *metrics.ThirdPartyMetrics
}

func (r repository) Create(ctx echo.Context, thread domain.Thread) error {
	outerSource := "http://vk-golang.ru:15000/thread?stable=true"
	source := "Create"
	subject := "thread"

	reqBody, err := json.Marshal(thread)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, outerSource, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

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
		return err
	}

	ctx.Logger().Infof("[%s] req_id:%d domain:%s src:%s endpoint:%s timing:%s",
		start.Format("2006-01-02 15:04:05"),
		ctx.Get(middleware.RequestIDKey).(uint64),
		subject,
		source,
		outerSource,
		timing.String(),
	)

	r.mtr.Timings.WithLabelValues(outerSource).Observe(timing.Seconds())
	r.mtr.Hits.WithLabelValues(resp.Status, outerSource).Inc()

	if resp.StatusCode != 200 {
		return errors.New("failed to create thread remotely")
	}

	return nil
}

func (r repository) Get(ctx echo.Context, id string) (domain.Thread, error) {
	outerSource := fmt.Sprintf("http://vk-golang.ru:15000/thread?id=%s&thread_fast=true", id)
	source := "Get"
	subject := "thread"

	req, err := http.NewRequest(http.MethodGet, outerSource, nil)
	if err != nil {
		return domain.Thread{}, err
	}

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
		return domain.Thread{}, err
	}

	ctx.Logger().Infof("[%s] req_id:%d domain:%s src:%s endpoint:%s timing:%s",
		start.Format("2006-01-02 15:04:05"),
		ctx.Get(middleware.RequestIDKey).(uint64),
		subject,
		source,
		outerSource,
		timing.String(),
	)

	r.mtr.Timings.WithLabelValues(outerSource).Observe(timing.Seconds())
	r.mtr.Hits.WithLabelValues(resp.Status, outerSource).Inc()

	if resp.StatusCode != 200 {
		return domain.Thread{}, errors.New("failed to fetch thread remotely")
	}

	var thread domain.Thread
	err = json.NewDecoder(resp.Body).Decode(&thread)
	if err != nil {
		return domain.Thread{}, err
	}

	return thread, nil
}

func NewRepository(m *metrics.ThirdPartyMetrics) domain.ThreadRepository {
	return repository{
		mtr: m,
	}
}
