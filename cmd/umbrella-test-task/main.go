package main

import (
	"middleware/internal/pkg/app"

	"github.com/rs/zerolog/log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Error().Err(err)
	}
	err = a.Run()
	if err != nil {
		log.Error().Err(err)
	}
}
