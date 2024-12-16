package domain

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

var ErrNoSession = errors.New("no session")

type Session struct {
	UserID string
}

type SessionService interface {
	CheckSession(ctx echo.Context, headers http.Header) (Session, error)
}
