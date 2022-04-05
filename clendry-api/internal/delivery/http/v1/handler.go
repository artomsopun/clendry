package v1

import (
	"github.com/artomsopun/clendry/clendry-api/internal/service"
	"github.com/artomsopun/clendry/clendry-api/pkg/auth"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
	filesManager files.Files
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager, filesManager files.Files) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		filesManager: filesManager,
	}
}

func (h *Handler) Init(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
	}
}
