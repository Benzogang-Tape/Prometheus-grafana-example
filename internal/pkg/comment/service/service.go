package service

import (
	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
)

type service struct {
	CommentRepo domain.CommentRepository
	ThreadRepo  domain.ThreadRepository
}

func NewService(commentRepo domain.CommentRepository, threadRepo domain.ThreadRepository) domain.CommentService {
	return service{
		CommentRepo: commentRepo,
		ThreadRepo:  threadRepo,
	}
}

func (s service) Create(ctx echo.Context, threadID string, comment domain.Comment) error {
	if err := s.checkThread(ctx, threadID); err != nil {
		return err
	}

	return s.CommentRepo.Create(ctx, comment)
}

func (s service) Like(ctx echo.Context, threadID string, commentID string) error {
	if err := s.checkThread(ctx, threadID); err != nil {
		return err
	}

	return s.CommentRepo.Like(ctx, commentID)
}

func (s service) checkThread(ctx echo.Context, threadID string) error {
	_, err := s.ThreadRepo.Get(ctx, threadID)
	return err
}
