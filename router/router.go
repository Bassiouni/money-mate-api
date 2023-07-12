package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RootRouter = gin.Default()

func BindRootRouter() {
	RootRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome!!!",
		})
	})
}
