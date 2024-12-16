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

func NewRepository(m *metrics.ThirdPartyMetrics) domain.CommentRepository {
	return repository{
		mtr: m,
	}
}

func (r repository) Create(ctx echo.Context, comment domain.Comment) error {
	outerSource := "http://vk-golang.ru:16000/comment?fast=true"
	source := "Create"
	subject := "comment"

	reqBody, err := json.Marshal(comment)
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
		return errors.New("failed to create comment remotely")
	}

	return nil
}

func (r repository) Like(ctx echo.Context, commentID string) error {
	outerSource := fmt.Sprintf("http://vk-golang.ru:16000/comment/like?cid=%s&superstable=true", commentID)
	source := "Like"
	subject := "comment"

	req, err := http.NewRequest(
		http.MethodPost,
		outerSource,
		nil,
	)
	if err != nil {
		return err
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
		return errors.New("failed to like comment remotely")
	}

	return nil
}
