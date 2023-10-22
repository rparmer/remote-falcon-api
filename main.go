package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/config"
	"github.com/rparmer/remote-falcon-api/internal/database"
	"github.com/rparmer/remote-falcon-api/internal/middleware"
	"github.com/rparmer/remote-falcon-api/internal/routes"
	"github.com/rparmer/remote-falcon-api/internal/service"
	"github.com/rparmer/remote-falcon-api/internal/util"
)

func main() {
	logger := util.NewLogger()
	c := config.GetConfig()
	mode := gin.ReleaseMode

	if c.Server.Debug {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	router := gin.New()
	router.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))

	db, _ := database.New()
	svc := service.New(db)
	middleware.Configure(db)

	apiRouter := router.Group("/")
	routes.AddRoutes(apiRouter, svc)

	port := fmt.Sprintf(":%d", c.Server.Port)
	logger.Info(fmt.Sprintf("listening on port %v", port))
	err := router.Run(port)
	if err != nil {
		logger.Fatal(err)
	}
}

func createServices(db *database.MongoDB) {

}
