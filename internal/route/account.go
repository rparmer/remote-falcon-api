package route

import "github.com/gin-gonic/gin"

func accountRoutes(superRoute *gin.RouterGroup) {
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
