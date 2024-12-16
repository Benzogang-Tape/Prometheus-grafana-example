package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/api/middleware"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
)

type Handler struct {
	ThreadSvc domain.ThreadService
}

func (h Handler) GetThread(ctx echo.Context) error {
	tid := ctx.Param("tid")
	source := "GetThread"
	subject := "thread"

	start := time.Now()
	t, err := h.ThreadSvc.Get(ctx, tid)
	timing := time.Since(start)
	if err != nil {
		ctx.Logger().Errorf("[%s] req_id:%d user_id:%s domain:%s src:%s timing:%s err:%s",
			time.Now().Format("2006-01-02 15:04:05"),
			ctx.Get(middleware.RequestIDKey).(uint64),
			ctx.Get(middleware.SessionKey).(domain.Session).UserID,
			subject,
			source,
			timing.String(),
			err.Error(),
		)
		ctx.NoContent(http.StatusInternalServerError) //nolint:errcheck
		return err
	}
	ctx.Logger().Infof("[%s] req_id:%d user_id:%s domain:%s src:%s timing:%s",
		time.Now().Format("2006-01-02 15:04:05"),
		ctx.Get(middleware.RequestIDKey).(uint64),
		ctx.Get(middleware.SessionKey).(domain.Session),
		subject,
		source,
		timing.String(),
	)

	return ctx.JSON(200, t)
}

func (h Handler) CreateThread(ctx echo.Context) error {
	source := "CreateThread"
	subject := "thread"
	var thread domain.Thread

	err := ctx.Bind(&thread)
	if err != nil {
		ctx.Logger().Errorf("[%s] req_id:%d user_id:%s domain:%s src:%s err:%s",
			time.Now().Format("2006-01-02 15:04:05"),
			ctx.Get(middleware.RequestIDKey).(uint64),
			ctx.Get(middleware.SessionKey).(domain.Session),
			subject,
			source,
			err.Error(),
		)
		ctx.NoContent(http.StatusBadRequest) //nolint:errcheck
		return err
	}

	start := time.Now()
	err = h.ThreadSvc.Create(ctx, thread)
	timing := time.Since(start)
	if err != nil {
		ctx.Logger().Errorf("[%s] req_id:%d user_id:%s domain:%s src:%s timing:%s err:%s",
			time.Now().Format("2006-01-02 15:04:05"),
			ctx.Get(middleware.RequestIDKey).(uint64),
			ctx.Get(middleware.SessionKey).(domain.Session).UserID,
			subject,
			source,
			timing.String(),
			err.Error(),
		)
		ctx.NoContent(http.StatusInternalServerError) //nolint:errcheck
		return err
	}
	ctx.Logger().Infof("[%s] req_id:%d user_id:%s domain:%s src:%s timing:%s",
		time.Now().Format("2006-01-02 15:04:05"),
		ctx.Get(middleware.RequestIDKey).(uint64),
		ctx.Get(middleware.SessionKey).(domain.Session),
		subject,
		source,
		timing.String(),
	)

	return ctx.NoContent(200)
}
