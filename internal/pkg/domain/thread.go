package domain

import (
	"github.com/labstack/echo"
)

type Thread struct {
	ID   string
	Name string
}

type ThreadService interface {
	Create(ctx echo.Context, thread Thread) error
	Get(ctx echo.Context, id string) (Thread, error)
}

type ThreadRepository interface {
	Create(ctx echo.Context, thread Thread) error
	Get(ctx echo.Context, id string) (Thread, error)
}
