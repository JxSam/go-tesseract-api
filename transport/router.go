package transport

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.RouterGroup, h *Handler) {
	rg := r.Group("/")
	{
		rg.GET("/health", h.Health)
	}
}
