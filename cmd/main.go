package main

import (
	"go-tesseract-api-docker/internal/app"
	"go-tesseract-api-docker/internal/config"
	"go-tesseract-api-docker/pkg/configloader"
)

func main() {
	cfg, err := configloader.Load[config.Config]()
	if err != nil {
		panic(err)
	}

	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}
