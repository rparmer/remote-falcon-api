package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/service"
)

func accountRoutes(superRoute *gin.RouterGroup, svc *service.Service) {
	accountRouter := superRoute.Group("/account")
	{
		// accountRouter.JSON
		accountRouter.GET("/ip", func(c *gin.Context) {
			c.String(200, "hello from ip")
		})

		// accountRouter.POST("/signUp" /*todo*/)
		// accountRouter.POST("/signIn" /*todo*/)
		// accountRouter.POST("/requestResetPassword" /*todo*/)
		// accountRouter.POST("/forgotPassword" /*todo*/)
		// accountRouter.POST("/resendVerificationEmail" /*todo*/)
		// accountRouter.POST("/verifyEmail" /*todo*/)
		// accountRouter.POST("/verifyResetPasswordLink" /*todo*/)
		// accountRouter.POST("/checkRemoteTokenExists" /*todo*/)
		// accountRouter.POST("/resetPassword" /*todo*/)
	}
}
