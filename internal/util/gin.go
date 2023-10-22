package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "404 page not found")
}

func GinServerError(c *gin.Context) {
	c.String(http.StatusInternalServerError, "500 internal server error")
}

func GinConflict(c *gin.Context) {
	c.String(http.StatusConflict, "409 duplicate key")
}

func GinOkString(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func GinOkJson(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}
