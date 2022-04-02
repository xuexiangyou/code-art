package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/common"
)

func NewRouter(handler *gin.Engine, f common.FxCommonParams) {
	// Routers
	h := handler.Group("/v1")
	{
		newTagRoutes(h, f)
		newArticleRoutes(h, f)
	}
}
