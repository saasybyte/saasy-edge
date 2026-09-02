package provider_model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// ListProviderModels implements api.ServerInterface
func (h *Handler) ListProviderModels(c *gin.Context) {
	models, err := h.service.ListProviderModels(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toAPIProviderModelsResponse(models))
}
