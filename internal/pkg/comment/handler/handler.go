package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/api/middleware"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
)

type Handler struct {
	CommentSvc domain.CommentService
}

func (h Handler) Create(ctx echo.Context) error {
	source := "Create"
	subject := "comment"
	var comment domain.Comment

	err := ctx.Bind(&comment)
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

	tid := ctx.Param("tid")

	start := time.Now()
	err = h.CommentSvc.Create(ctx, tid, comment)
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

	return nil
}

func (h Handler) Like(ctx echo.Context) error {
	source := "Like"
	subject := "comment"
	tid := ctx.Param("tid")
	cid := ctx.Param("cid")

	start := time.Now()
	err := h.CommentSvc.Like(ctx, tid, cid)
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

	return nil
}
