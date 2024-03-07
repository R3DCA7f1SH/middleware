package app

import (
	"fmt"
	"log"
	"middleware/internal/app/endpoint"
	"middleware/internal/app/mw"
	"middleware/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.GET("/status", a.e.Status, mw.RoleCheck)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server runnung...")

	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal("ERRRORRR!!!!")
	}

	return nil
}