package tesseract

import (
	"go-tesseract-api-docker/pkg/logger"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) GetText(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file is required",
		})
		return
	}

	tempDir, err := os.MkdirTemp("", "ocr-*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer os.RemoveAll(tempDir)

	pdfPath := filepath.Join(tempDir, file.Filename)

	if err := c.SaveUploadedFile(file, pdfPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	text, err := GetText(pdfPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text": *text,
	})
}
