package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Hello world",
		})
	})
	r.Run(port)
}
