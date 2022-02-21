package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/controllers"
	"go.uber.org/dig"
)

type RouterParams struct {
	dig.In
	TagController *controllers.TagController
}


func InitRouter(p RouterParams) *gin.Engine {
	r := gin.Default()

	r.GET("/get-tag", p.TagController.GetTag)

	return r
}
