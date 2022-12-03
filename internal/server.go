package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/ping", ping)

	router.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
