package tesseract

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.RouterGroup, h *Handler) {
	rg := r.Group("/pdf")
	{
		rg.POST("/to_text", h.GetText)
	}
}
