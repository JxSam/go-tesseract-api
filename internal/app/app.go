package app

import (
	"fmt"
	"go-tesseract-api-docker/internal/config"
	"go-tesseract-api-docker/internal/tesseract"
	"go-tesseract-api-docker/pkg/httpserver"
	"go-tesseract-api-docker/pkg/logger"
	"go-tesseract-api-docker/transport"
)

func Run(cfg *config.Config) error {
	logger := logger.New()

	httpserver := httpserver.New()

	api := httpserver.Group("/tesseract")

	handler := transport.NewHandler()

	transport.SetupRouter(api, handler)

	tesseract.NewModule(api, logger)

	addr := fmt.Sprintf("%s:%d", cfg.ServerConfig.Host, cfg.ServerConfig.Port)
	httpserver.Run(addr)

	return nil
}
