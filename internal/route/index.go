package route

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	accountRoutes(superRoute)
}
