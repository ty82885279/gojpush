package routers

import (
	"gojpush/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	r := gin.Default()
	r.GET("/ios", controller.SendMsg)

	return r
}
