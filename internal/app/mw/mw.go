package mw

import (
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(val, roleAdmin) {
			log.Info().Msg("red button user detected")
		}

		err := next(ctx)

		if err != nil {
			return err
		}

		return nil
	}
}
