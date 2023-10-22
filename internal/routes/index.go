package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/service"
)

func AddRoutes(parentRoute *gin.RouterGroup, svc *service.Service) {
	accountRoutes(parentRoute, svc)
	// authRoutes(parentRoute, svc)
	pluginRoutes(parentRoute, svc)
	preferencesRoutes(parentRoute, svc)
	sequenceRoutes(parentRoute, svc)
	viewerRoutes(parentRoute, svc)
}
