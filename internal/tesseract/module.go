package tesseract

import (
	"go-tesseract-api-docker/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewModule(r *gin.RouterGroup, logger *logger.Logger) {

	Handler := NewHandler(logger)

	SetupRouter(r, Handler)
}
