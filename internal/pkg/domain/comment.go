package domain

import (
	"github.com/labstack/echo"
)

type Comment struct {
	ID   string
	Text string
}

type CommentService interface {
	Create(ctx echo.Context, threadID string, comment Comment) error
	Like(ctx echo.Context, threadID string, commentID string) error
}

type CommentRepository interface {
	Create(ctx echo.Context, comment Comment) error
	Like(ctx echo.Context, commentID string) error
}
