package route

import "github.com/gin-gonic/gin"

func SetUp(gin *gin.Engine) {
	publicRouter := gin.Group("api")
	TestRouter(publicRouter)
}
