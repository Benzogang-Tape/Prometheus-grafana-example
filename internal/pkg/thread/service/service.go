package service

import (
	"github.com/labstack/echo"

	"github.com/Benzogang-Tape/Prometheus-grafana-example/internal/pkg/domain"
)

type service struct {
	ThreadRepo domain.ThreadRepository
}

func NewService(threadRepo domain.ThreadRepository) domain.ThreadService {
	return service{
		ThreadRepo: threadRepo,
	}
}

func (s service) Create(ctx echo.Context, thread domain.Thread) error {
	return s.ThreadRepo.Create(ctx, thread)
}

func (s service) Get(ctx echo.Context, id string) (domain.Thread, error) {
	return s.ThreadRepo.Get(ctx, id)
}
