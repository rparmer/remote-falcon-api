package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/service"
	"github.com/rparmer/remote-falcon-api/internal/util"
)

func preferencesRoutes(superRoute *gin.RouterGroup, svc *service.Service) {
	logger := util.GetLogger()
	preferencesRouter := superRoute.Group("/preferences")

	preferencesRouter.GET("", func(c *gin.Context) {
		prefs, err := svc.GetPreferences()
		if err != nil {
			logger.Error("no preferences found")
			util.GinNotFound(c)
			return
		}
		util.GinOkJson(c, &prefs)
	})

	preferencesRouter.POST("", func(c *gin.Context) {
		var prefs service.Preferences
		if err := c.BindJSON(&prefs); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.CreatePreferences(prefs); err != nil {
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

	preferencesRouter.PUT("", func(c *gin.Context) {
		var prefs service.Preferences
		if err := c.BindJSON(&prefs); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.UpdatePreferences(prefs); err != nil {
			logger.Error(err)
			if strings.Contains(err.Error(), "duplicate key") {
				util.GinConflict(c)
				return
			}
			if strings.Contains(err.Error(), "no records") {
				util.GinNotFound(c)
				return
			}
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})

	preferencesRouter.DELETE("", func(c *gin.Context) {
		var prefs service.Preferences
		if err := c.BindJSON(&prefs); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.DeletePreferences(prefs); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})
}
