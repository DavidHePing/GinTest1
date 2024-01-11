package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestRouter(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "Test Get")
	})

	routeGroup.GET("/test/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, fmt.Sprint("Test Get id = ", id))
	})

	routeGroup.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "Test Post")
	})

	routeGroup.PATCH("/test/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, fmt.Sprint("Test Patch id = ", id))
	})

	routeGroup.DELETE("/test/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, fmt.Sprint("Test Delete id = ", id))
	})
}
