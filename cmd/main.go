package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/api/middleware"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/comment/handler"
	commentrepo "github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/comment/repository"
	commentsvc "github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/comment/service"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain/metrics"
	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/session"
	threadhttp "github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/thread/handler"
	threadrepo "github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/thread/repository"
	threadsvc "github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/thread/service"
)

func main() {
	e := echo.New()
	m := echo.New()
	e.Logger.SetLevel(log.ERROR)

	mtr := metrics.NewMetrics()
	sessionSvc := session.NewService(mtr.ThirdParty)

	e.Use(middleware.RequestIDMiddleware())
	e.Use(middleware.MetricsWriterMiddleware(mtr.Native))
	e.Use(middleware.AuthEchoMiddleware(sessionSvc))
	e.Use(middleware.AccessLogMiddleware())

	threadRepo := threadrepo.NewRepository(mtr.ThirdParty)
	threadSvc := threadsvc.NewService(threadRepo)
	threadHandler := threadhttp.Handler{ThreadSvc: threadSvc}

	commentRepo := commentrepo.NewRepository(mtr.ThirdParty)
	commentSvc := commentsvc.NewService(commentRepo, threadRepo)
	commentHandler := handler.Handler{CommentSvc: commentSvc}

	m.Any("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/thread/:id", threadHandler.GetThread)
	e.POST("/thread", threadHandler.CreateThread)
	e.POST("/thread/:tid/comment", commentHandler.Create)
	e.POST("/thread/:tid/comment/:cid/like", commentHandler.Like)

	go func() {
		fmt.Println(m.Start(":8081"))
	}()
	fmt.Print(e.Start(":8000"))
}
