package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/service"
	"github.com/rparmer/remote-falcon-api/internal/util"
)

func sequenceRoutes(superRoute *gin.RouterGroup, svc *service.Service) {
	logger := util.GetLogger()
	sequenceRouter := superRoute.Group("/sequence")

	// sequence routes
	sequenceRouter.GET("", func(c *gin.Context) {
		name := c.Query("name")
		sequence, err := svc.GetSequence(name)
		if err != nil {
			logger.Error("no sequence found")
			util.GinNotFound(c)
			return
		}
		util.GinOkJson(c, &sequence)
	})

	sequenceRouter.POST("", func(c *gin.Context) {
		var sequence service.Sequence
		if err := c.BindJSON(&sequence); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.CreateSequence(sequence); err != nil {
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

	sequenceRouter.PUT("", func(c *gin.Context) {
		var sequence service.Sequence
		if err := c.BindJSON(&sequence); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.UpdateSequence(sequence); err != nil {
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

	sequenceRouter.DELETE("", func(c *gin.Context) {
		var sequence service.Sequence
		if err := c.BindJSON(&sequence); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.DeleteSequence(sequence); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})

	// sequence group routes
	sequenceRouter.GET("/group", func(c *gin.Context) {
		name := c.Query("name")
		sequenceGroup, err := svc.GetSequenceGroup(name)
		if err != nil {
			logger.Error("no sequence found")
			util.GinNotFound(c)
			return
		}
		util.GinOkJson(c, &sequenceGroup)
	})

	sequenceRouter.POST("/group", func(c *gin.Context) {
		var sequenceGroup service.SequenceGroup
		if err := c.BindJSON(&sequenceGroup); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.CreateSequenceGroup(sequenceGroup); err != nil {
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

	sequenceRouter.PUT("/group", func(c *gin.Context) {
		var sequenceGroup service.SequenceGroup
		if err := c.BindJSON(&sequenceGroup); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.UpdateSequenceGroup(sequenceGroup); err != nil {
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

	sequenceRouter.DELETE("/group", func(c *gin.Context) {
		var sequenceGroup service.SequenceGroup
		if err := c.BindJSON(&sequenceGroup); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		if err := svc.DeleteSequenceGroup(sequenceGroup); err != nil {
			logger.Error(err)
			util.GinServerError(c)
			return
		}
		util.GinOkString(c)
	})
}
