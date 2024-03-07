package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		s: svc,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()

	str := fmt.Sprintf("Days left: %d", d)

	err := ctx.String(http.StatusOK, str)

	if err != nil {
		return err
	}

	return nil
}
