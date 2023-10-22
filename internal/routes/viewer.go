package routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/service"
	"github.com/rparmer/remote-falcon-api/internal/util"
)

func viewerRoutes(superRoute *gin.RouterGroup, svc *service.Service) {
	logger := util.GetLogger()
	viewerRoutes := superRoute.Group("/viewer")
	// auth := util.GetAuth()
	// viewerRoutes.use(auth.)
	// viewerRoutes.Use(middleware.TokenAuthRequired)

	// viewer page routes
	viewerRoutes.GET("/page", func(c *gin.Context) {
		name := c.Query("name")
		page, err := svc.GetViewerPage(name)
		if err != nil {
			logger.Error(fmt.Sprintf("viewer page not found: %s", name))
			util.GinNotFound(c)
			return
		}
		util.GinOkJson(c, &page)
	})

	viewerRoutes.GET("/pages", func(c *gin.Context) {
		pages, err := svc.ListViewerPages()
		if err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkJson(c, &pages)
	})

	viewerRoutes.POST("/page", func(c *gin.Context) {
		var page service.ViewerPage
		if err := c.BindJSON(&page); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.CreateViewerPage(page); err != nil {
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

	viewerRoutes.PUT("/page", func(c *gin.Context) {
		var page service.ViewerPage
		if err := c.BindJSON(&page); err != nil {
			logger.Error(err)
			return
		}
		if err := svc.UpdateViewerPage(page); err != nil {
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

	viewerRoutes.DELETE("/page", func(c *gin.Context) {
		var page service.ViewerPage
		if err := c.BindJSON(&page); err != nil {
			logger.Error(err)
			return
		}
		if err := svc.DeleteViewerPage(page); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})

	// viewer page template routes
	viewerRoutes.GET("/pageTemplate", func(c *gin.Context) {
		name := c.Query("name")
		page, err := svc.GetViewerPageTemplate(name)
		if err != nil {
			logger.Error(fmt.Sprintf("viewer page not found: %s", name))
			util.GinNotFound(c)
			return
		}
		util.GinOkJson(c, &page)
	})

	viewerRoutes.GET("/pageTemplates", func(c *gin.Context) {
		pages, err := svc.ListViewerPageTemplates()
		if err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkJson(c, &pages)
	})

	viewerRoutes.POST("/pageTemplate", func(c *gin.Context) {
		var page service.ViewerPageTemplate
		if err := c.BindJSON(&page); err != nil {
			logger.Error(err)
			return
		}
		if err := svc.CreateViewerPageTemplate(page); err != nil {
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

	viewerRoutes.DELETE("/pageTemplate", func(c *gin.Context) {
		var page service.ViewerPageTemplate
		if err := c.BindJSON(&page); err != nil {
			logger.Error(err)
			return
		}
		if err := svc.DeleteViewerPageTemplate(page); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})
}
