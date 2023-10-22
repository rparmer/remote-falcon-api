package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/service"
	"github.com/rparmer/remote-falcon-api/internal/util"
)

func pluginRoutes(superRoute *gin.RouterGroup, svc *service.Service) {
	logger := util.GetLogger()
	pluginRouter := superRoute.Group("/plugin")

	pluginRouter.GET("/queue", func(c *gin.Context) {
		c.String(200, "hello from ip")
	})

	pluginRouter.PUT("/queue", func(c *gin.Context) {
		c.String(200, "hello from ip")
	})

	pluginRouter.DELETE("/queue", func(c *gin.Context) {
		c.String(200, "hello from ip")
	})

	pluginRouter.POST("/sync", func(c *gin.Context) {
		var sequences []service.Sequence
		if err := c.BindJSON(&sequences); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.PluginSyncSequences(sequences); err != nil {
			logger.Error(err)
			if strings.Contains(err.Error(), "duplicate key") {
				util.GinConflict(c)
				return
			}
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})

}
